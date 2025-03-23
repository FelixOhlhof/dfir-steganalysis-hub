import React, { useState } from "react";
import axios from "axios";
import { Container } from "react-bootstrap";
import ReactJson from "react-json-view";
import { useAlert } from "react-alert";
import config from "../config";

function WorkflowDemo(props) {
  const [file, setFile] = useState(null);
  const [jsonData, setJsonData] = useState(null);

  const alert = useAlert();

  const analysisEndpoint = `${config.restGwUrl}/execute`;

  const handleFileChange = (event) => {
    setJsonData(null);
    setFile(event.target.files[0]);
  };

  const handleExecute = (event) => {
    event.preventDefault();

    const formData = new FormData();
    formData.append("file", file);

    const httpConfig = {
      headers: {
        "content-type": "multipart/form-data",
      },
    };

    axios
      .post(analysisEndpoint, formData, httpConfig)
      .then((response) => {
        setJsonData(response.data);
        console.debug(response.data);
      })
      .catch((error) => {
        console.error("Error uploading files: ", error);
        alert.show(`Error uploading files: ${error} ${error?.response?.data}`, {
          type: "error",
        });
      });
  };

  return (
    <Container className="my-5">
      <h2 className="">Execute Workflow</h2>
      <form>
        <div className="input-group">
          <input
            className="form-control"
            type="file"
            id="formFile"
            onChange={handleFileChange}
          />
          <button
            className="btn btn btn-primary"
            type="submit"
            onClick={handleExecute}
            id="button-addon2"
          >
            Execute
          </button>
        </div>
      </form>
      {jsonData ? (
        <div className="container border rounded">
          <ReactJson
            className="border"
            src={jsonData}
            theme="summerfruit:inverted"
            style={{
              whiteSpace: "pre-wrap", // Erzwingt Zeilenumbruch
              wordBreak: "break-word", // Lange Wörter umbrechen
            }}
          />
        </div>
      ) : (
        <div />
      )}
    </Container>
  );
}

export default WorkflowDemo;
