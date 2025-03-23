#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>
#include <cjson/cJSON.h>
#include <openssl/sha.h>
#include <yara/modules.h>
#include <time.h>


#define MODULE_NAME stego
#define DICT_NAME "values"
#define MAX_PARAMS 5
#define SHA256_STRING_LENGTH (SHA256_DIGEST_LENGTH * 2 + 1)
#define LOG_FILE "log.txt"

CRITICAL_SECTION log_mutex;
FILE* log_file = NULL;
int errors = 0;
int successfull = 0;

struct File
{
	size_t file_size;
	char* file_data;
	char* file_name;
};

struct ResponseData
{
	char* data;
	size_t size;
};

typedef struct {
	char* buffer;
	size_t size;
} LogBuffer;


void init_logger() {
	InitializeCriticalSection(&log_mutex);
	log_file = fopen(LOG_FILE, "a");
	if (log_file == NULL) {
		perror("error opening log file");
		exit(EXIT_FAILURE);
	}
}

void log_message(const char* format, ...) {
	if (log_file == NULL) return;

	EnterCriticalSection(&log_mutex);


	va_list args;
	va_start(args, format);
	size_t message_len = vsnprintf(NULL, 0, format, args) + 1;
	va_end(args);

	if (message_len <= 0) {
		LeaveCriticalSection(&log_mutex);
		return;
	}

	char* message = (char*)malloc(message_len);
	if (message == NULL) {
		perror("malloc failed");
		LeaveCriticalSection(&log_mutex);
		exit(EXIT_FAILURE);
	}

	va_start(args, format);
	vsnprintf(message, message_len, format, args);
	va_end(args);

	fprintf(log_file, "%s\n", message);
	fflush(log_file);

	free(message);
	LeaveCriticalSection(&log_mutex);
}

void close_logger() {
	if (log_file != NULL) {
		fclose(log_file);
		log_file = NULL;
	}
	DeleteCriticalSection(&log_mutex);
}

void init_log_buffer(LogBuffer* log) {
	log->buffer = NULL;
	log->size = 0;
}

void add_log_msg(LogBuffer* log, const char* format, ...) {
	if (log == NULL || format == NULL) return;

	va_list args;
	va_start(args, format);

	size_t message_len = vsnprintf(NULL, 0, format, args) + 1;
	va_end(args);

	if (message_len <= 0) return;

	char* temp_msg = (char*)malloc(message_len);
	if (temp_msg == NULL) {
		perror("malloc failed");
		exit(EXIT_FAILURE);
	}

	va_start(args, format);
	vsnprintf(temp_msg, message_len, format, args);
	va_end(args);

	size_t new_size = log->size + message_len;
	char* temp_buffer = (char*)realloc(log->buffer, new_size);
	if (temp_buffer == NULL) {
		perror("realloc failed");
		free(temp_msg);
		exit(EXIT_FAILURE);
	}

	log->buffer = temp_buffer;
	if (log->size == 0) {
		strcpy(log->buffer, temp_msg);
	}
	else {
		strcat(log->buffer, temp_msg);
	}

	log->size = new_size - 1;
	free(temp_msg);
}

void free_log_buffer(LogBuffer* log) {
	free(log->buffer);
	log->buffer = NULL;
	log->size = 0;
}

void calculate_sha256_from_buffer(const char* file_data, size_t file_size, unsigned char* output_hash) {
	SHA256_CTX sha256;
	SHA256_Init(&sha256);
	SHA256_Update(&sha256, file_data, file_size);
	SHA256_Final(output_hash, &sha256);
}

void sha256_to_hex(const unsigned char* hash, char* output_hex) {
	for (int i = 0; i < SHA256_DIGEST_LENGTH; i++) {
		sprintf(output_hex + (i * 2), "%02x", hash[i]);
	}
	output_hex[SHA256_STRING_LENGTH - 1] = '\0';
}

size_t write_callback(void* contents, size_t size, size_t nmemb, void* userp)
{
	size_t totalSize = size * nmemb;
	struct ResponseData* resp = (struct ResponseData*)userp;

	char* ptr = realloc(resp->data, resp->size + totalSize + 1);
	if (ptr == NULL)
	{
		errors++;
		return ERROR_ERRORS_ENCOUNTERED;
	}

	resp->data = ptr;
	memcpy(&(resp->data[resp->size]), contents, totalSize);
	resp->size += totalSize;
	resp->data[resp->size] = 0;

	return totalSize;
}

void parse_json_value_item(char* key, cJSON* value, YR_OBJECT* module_object, bool is_dict, LogBuffer* log) {
	if (value->child != NULL) {
		bool dict = strcmp(value->string, "StructuredValue") == 0;
		parse_json_value_item(key, value->child, module_object, dict, log);
		return;
	}
	if (value->next != NULL) {
		parse_json_value_item(key, value->next, module_object, is_dict, log);
	}

	char* newKey = key;
	/*if (is_dict) {
		size_t len = strlen(key) + strlen(value->string) + 2;
		newKey = (char*)malloc(len);
		snprintf(newKey, len, "%s.%s", key, value->string);
	}*/

	char* parsed_value = "";

	switch (value->type)
	{
	case cJSON_String:
		parsed_value = value->valuestring;
		break;
	case cJSON_Number:
		sprintf(parsed_value, "%f", value->valuedouble);
		break;
	case cJSON_True:
		parsed_value = "true";
		break;
	case cJSON_False:
		parsed_value = "false";
		break;
	}

	char key_combination[256];
	snprintf(key_combination, sizeof(key_combination), "%s[\"%s\"]", DICT_NAME, newKey);
	yr_set_string(parsed_value, module_object, key_combination);

	add_log_msg(log, newKey);
	add_log_msg(log, ": ");
	add_log_msg(log, parsed_value);
	add_log_msg(log, "\n");
}

void parse_and_set_values(const char* json_str, YR_OBJECT* module_object, LogBuffer* log) {
	cJSON* root = cJSON_Parse(json_str);
	if (!root) {
		errors++;
		fprintf(stderr, "error while parsing json\n");
		add_log_msg(&log, "error while parsing json\n");
		return;
	}

	cJSON* task_results = cJSON_GetObjectItem(root, "task_results");
	if (!cJSON_IsArray(task_results)) {
		cJSON_Delete(root);
		return;
	}

	cJSON* task_result;
	cJSON_ArrayForEach(task_result, task_results) {
		cJSON* task_id = cJSON_GetObjectItem(task_result, "task_id");
		cJSON* service_response = cJSON_GetObjectItem(task_result, "service_response");
		cJSON* error = cJSON_GetObjectItem(service_response, "error");

		if (!cJSON_IsString(task_id)) {
			continue;
		}

		if (cJSON_IsObject(service_response)) {
			cJSON* values = cJSON_GetObjectItem(service_response, "values");
			if (cJSON_IsObject(values)) {
				cJSON* value_item;
				cJSON_ArrayForEach(value_item, values) {
					char key_combination[256];
					snprintf(key_combination, sizeof(key_combination), "%s->%s", task_id->valuestring, value_item->string);

					parse_json_value_item(key_combination, value_item, module_object, false, log);

				}
			}
		}

		if (cJSON_IsString(error)) {
			fprintf(stderr, "%s\n", error->valuestring);

			char key_combination[256];
			snprintf(key_combination, sizeof(key_combination), "%s[\"%s->error\"]", DICT_NAME, task_id->valuestring);
			yr_set_string(error->valuestring, module_object, key_combination);
		}
	}

	cJSON_Delete(root);
}


int send_file_data(char* endpoint, struct File* f, char** params, char* service_name, struct ResponseData* response, LogBuffer* log) {
	unsigned char hash[SHA256_DIGEST_LENGTH];
	calculate_sha256_from_buffer(f->file_data, f->file_size, hash);
	char hash_hex[SHA256_STRING_LENGTH];
	sha256_to_hex(hash, hash_hex);
	add_log_msg(log, "File: %s\n", hash_hex);

	CURL* curl;
	CURLcode res;

	curl_global_init(CURL_GLOBAL_DEFAULT);
	curl = curl_easy_init();

	if (curl)
	{
		curl_mime* mime;
		curl_mimepart* part;

		mime = curl_mime_init(curl);

		part = curl_mime_addpart(mime);
		curl_mime_name(part, "file");
		curl_mime_filename(part, "unknown"); // yara doesn't include the file name anymore
		curl_mime_data(part, f->file_data, f->file_size);

		for (int i = 0; i <= MAX_PARAMS; i++) {
			char* p = params[i];

			if (p == NULL) {
				break;
			}

			char* key = strtok(p, ":");
			char* value = strtok(NULL, ":");

			part = curl_mime_addpart(mime);
			curl_mime_name(part, key);
			curl_mime_data(part, value, strlen(value));
		}

		if (service_name != NULL) {
			part = curl_mime_addpart(mime);
			curl_mime_name(part, "exec");
			curl_mime_data(part, service_name, strlen(service_name));
		}

		curl_easy_setopt(curl, CURLOPT_URL, endpoint);
		curl_easy_setopt(curl, CURLOPT_MIMEPOST, mime);
		curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
		curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void*)response);

		res = curl_easy_perform(curl);

		if (res != CURLE_OK)
		{
			errors++;
			return ERROR_ERRORS_ENCOUNTERED;
		}

		curl_mime_free(mime);
		curl_easy_cleanup(curl);
	}

	curl_global_cleanup();

	return res == CURLE_OK ? 0 : 1;
}

struct File* get_file_data(YR_SCAN_CONTEXT* context)
{
	YR_MEMORY_BLOCK* block;
	char* file_data = NULL;
	size_t file_size = 0;

	foreach_memory_block(context->iterator, block)
	{
		file_data = block->fetch_data(block);
		file_size = block->size;
		break;
	}

	struct File* f = malloc(sizeof(struct File));
	if (f == NULL)
	{
		errors++;
		return ERROR_ERRORS_ENCOUNTERED;
	}

	f->file_data = file_data;
	f->file_size = file_size;

	return f;
}

/// <summary>
/// Executes a specific service via the rest gateway
/// </summary>
/// <param name="service_name"></param>
/// <returns>Return Values of Service</returns>
char* execute_service(YR_OBJECT* module_object, char* endpoint, char* service_name, char* return_value_name, struct File* file, char** params, LogBuffer* log) {
	struct ResponseData response = { 0 };
	response.data = malloc(1);
	response.size = 0;

	if (send_file_data(endpoint, file, params, service_name, &response, log) == 0)
	{
		// success
		parse_and_set_values(response.data, module_object, log);
	}

	free(response.data);
	free(file);

	char key[256];
	snprintf(key, sizeof(key), "%s[\"%s->%s\"]", DICT_NAME, service_name, return_value_name);
	return yr_get_string(module_object, key)->c_string;
}

char** extract_params(YR_VALUE* __args, int offset) {
	char** params = malloc(MAX_PARAMS * sizeof(char*));
	if (params == NULL) {
		return NULL;
	}

	for (int i = 0; i < MAX_PARAMS; i++) {
		const char* value_name = string_argument(i + offset);

		if (value_name == 0xccccccccccccccd4) {
			params[i] = NULL;
			break;
		}

		params[i] = strdup(value_name);

		if (params[i] == NULL) {
			for (int j = 0; j < i; j++) {
				free(params[j]);
			}
			free(params);
			return NULL;
		}
	}

	return params;
}

define_function(exec) {
	YR_SCAN_CONTEXT* context = yr_scan_context();
	YR_OBJECT* module_object = yr_module();

	LogBuffer log;
	init_log_buffer(&log);

	char* endpoint = string_argument(1);
	char** params = extract_params(__args, 2);

	struct File* file = get_file_data(context);

	struct ResponseData response = { 0 };
	response.data = malloc(1);
	response.size = 0;

	if (send_file_data(endpoint, file, params, NULL, &response, &log) == 0) {
		parse_and_set_values(response.data, module_object, &log);
		successfull++;
	}
	else {
		errors++;
	}

	add_log_msg(&log, "-------------------------------------------------------------------\n");
	log_message(log.buffer);

	free_log_buffer(&log);
	free(response.data);
	free(file);

	return_integer(ERROR_SUCCESS);
}

define_function(i_exec)
{
	YR_SCAN_CONTEXT* context = yr_scan_context();
	YR_OBJECT* module_object = yr_module();

	LogBuffer log;
	init_log_buffer(&log);

	char* endpoint = string_argument(1);
	char* service_name = string_argument(2);
	char* return_value_name = string_argument(3);
	char** params = extract_params(__args, 4);

	struct File* file = get_file_data(context);

	char* return_value = execute_service(module_object, endpoint, service_name, return_value_name, file, params, &log);

	if (return_value == 0x0000000000000008) {
		errors++;
		fprintf(stderr, "%s not existing\n", return_value_name);
		add_log_msg(&log, "%s not existing\n", return_value_name);
		return_integer(YR_UNDEFINED);
	}
	else {
		successfull++;
	}

	add_log_msg(&log, "-------------------------------------------------------------------\n");
	log_message(log.buffer);

	return_integer(atoi(return_value));
}

define_function(f_exec)
{
	YR_SCAN_CONTEXT* context = yr_scan_context();
	YR_OBJECT* module_object = yr_module();

	LogBuffer log;
	init_log_buffer(&log);

	char* endpoint = string_argument(1);
	char* service_name = string_argument(2);
	char* return_value_name = string_argument(3);
	char** params = extract_params(__args, 4);

	struct File* file = get_file_data(context);

	char* return_value = execute_service(module_object, endpoint, service_name, return_value_name, file, params, &log);

	if (return_value == 0x0000000000000008) {
		errors++;
		fprintf(stderr, "%s not existing\n", return_value_name);
		add_log_msg(&log, "%s not existing\n", return_value_name);
		return_float(YR_UNDEFINED);
	}
	else {
		successfull++;
	}

	add_log_msg(&log, "-------------------------------------------------------------------\n");
	log_message(log.buffer);
	return_float(atof(return_value));
}

define_function(s_exec)
{
	YR_SCAN_CONTEXT* context = yr_scan_context();
	YR_OBJECT* module_object = yr_module();

	LogBuffer log;
	init_log_buffer(&log);

	char* endpoint = string_argument(1);
	char* service_name = string_argument(2);
	char* return_value_name = string_argument(3);
	char** params = extract_params(__args, 4);

	struct File* file = get_file_data(context);

	char* return_value = execute_service(module_object, endpoint, service_name, return_value_name, file, params, &log);

	if (return_value == 0x0000000000000008) {
		errors++;
		fprintf(stderr, "%s not existing\n", return_value_name);
		add_log_msg(&log, "%s not existing\n", return_value_name);
		return_string(YR_UNDEFINED);
	}
	else {
		successfull++;
	}

	add_log_msg(&log, "-------------------------------------------------------------------\n");
	log_message(log.buffer);

	return_string(return_value);
}

begin_declarations
declare_string_dictionary("values");

declare_function("exec", "s", "i", exec);

declare_function("s_exec", "ss", "s", s_exec);
declare_function("s_exec", "sss", "s", s_exec);
declare_function("s_exec", "ssss", "s", s_exec);
declare_function("s_exec", "sssss", "s", s_exec);
declare_function("s_exec", "ssssss", "s", s_exec);
declare_function("s_exec", "sssssss", "s", s_exec);

declare_function("i_exec", "ss", "i", i_exec);
declare_function("i_exec", "sss", "i", i_exec);
declare_function("i_exec", "ssss", "i", i_exec);
declare_function("i_exec", "sssss", "i", i_exec);
declare_function("i_exec", "ssssss", "i", i_exec);
declare_function("i_exec", "sssssss", "i", i_exec);

declare_function("f_exec", "ss", "f", f_exec);
declare_function("f_exec", "sss", "f", f_exec);
declare_function("f_exec", "ssss", "f", f_exec);
declare_function("f_exec", "sssss", "f", f_exec);
declare_function("f_exec", "ssssss", "f", f_exec);
declare_function("f_exec", "sssssss", "f", f_exec);
end_declarations


int module_initialize(YR_MODULE* module)
{
	init_logger();
	time_t now = time(NULL);
	struct tm time_info;
	localtime_s(&time_info, &now);

	char time_str[20];
	strftime(time_str, sizeof(time_str), "%Y-%m-%d %H:%M:%S", &time_info);
	log_message("Executing analysis: %s\n", time_str);
	return ERROR_SUCCESS;
}

int module_finalize(YR_MODULE* module)
{
	log_message("Successfull: %d  Errors: %d", successfull, errors);
	close_logger();
	return ERROR_SUCCESS;
}

int module_load(
	YR_SCAN_CONTEXT* context,
	YR_OBJECT* module_object,
	void* module_data,
	size_t module_data_size)
{
	return ERROR_SUCCESS;
}

int module_unload(YR_OBJECT* module_object)
{
	return ERROR_SUCCESS;
}
