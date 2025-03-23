import React, { useEffect, useState } from "react";
import axios from "axios";
import {
  Container,
  Popover,
  OverlayTrigger,
  Button,
  Badge,
} from "react-bootstrap";
import AceEditor from "react-ace";
import { Buffer } from "buffer";
import { useAlert } from "react-alert";
import "ace-builds/src-noconflict/theme-chrome";
import "ace-builds/src-noconflict/mode-yaml";
import config from "../config";

function WorkflowEditor(props) {
  const [workflow, setWorkflow] = useState("");
  const [workflowDemoYaml, setYamlData] = useState("");

  const alert = useAlert();
  const workflowEndpoint = `${config.restGwUrl}/workflow`;

  useEffect(() => {
    axios
      .get(workflowEndpoint)
      .then((response) => {
        const decodedFile = Buffer.from(response.data, "base64").toString(
          "utf-8"
        );
        setWorkflow(decodedFile);
      })
      .catch((error) => {
        console.error("Error fetching workflow:", error);
        alert.show(
          `Error fetching workflow: ${error} ${error?.response?.data}`,
          { type: "error" }
        );
      });
  }, []);

  useEffect(() => {
    fetch("/workflow_demo.yaml")
      .then((response) => response.text())
      .then((data) => {
        setYamlData(data);
      })
      .catch((error) => {
        console.error("error loading demo yaml:", error);
      });
  }, []);

  // Handle saving edited workflow.yaml
  const handleSaveWorkflow = () => {
    const encodedFile = Buffer.from(workflow, "utf-8").toString("base64");
    const httpConfig = {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
    };
    axios
      .post(workflowEndpoint, { file: encodedFile }, httpConfig)
      .then((response) => alert.show("Workflow updated", { type: "success" }))
      .catch((error) => {
        console.error("Error saving workflow:", error);
        alert.show(
          `Error updating workflow: ${error} ${error?.response?.data}`,
          {
            type: "error",
          }
        );
      });
  };

  // Handle load workflow.yaml
  const handleLoadWorkflow = (event) => {
    const file = event.target.files[0];
    if (!file) return;

    if (
      file.type !== "application/x-yaml" &&
      !file.name.endsWith(".yaml") &&
      !file.name.endsWith(".yml")
    ) {
      alert.show(`please choose a valid yaml file`, {
        type: "error",
      });
      return;
    }

    const reader = new FileReader();
    reader.onload = (e) => {
      try {
        setWorkflow(e.target.result);
      } catch (err) {
        alert.show(`error while parsing workflow file`, {
          type: "error",
        });
        console.error(err);
      }
    };

    reader.onerror = () => {
      alert.show(`error while reading workflow file`, {
        type: "error",
      });
    };

    reader.readAsText(file);
    event.target.value = "";
  };

  const popover = (
    <Popover id="custom-popover" style={{ "--bs-popover-max-width": "57em" }}>
      <Popover.Header>Example Workflow</Popover.Header>
      <Popover.Body>
        <AceEditor
          mode="yaml"
          theme="chrome"
          value={workflowDemoYaml}
          setOptions={{ useWorker: false }}
          name="workflow-editor"
          editorProps={{ $blockScrolling: true }}
          width="57em"
          height="25em"
          readOnly
        />
      </Popover.Body>
    </Popover>
  );

  const triggerFileInput = () => {
    document.getElementById("file-input").click();
  };

  return (
    <Container>
      <div className="row">
        <div className="col-8">
          <div
            style={{
              display: "flex",
              alignItems: "center",
            }}
          >
            <h2>Edit Workflow </h2>
            <OverlayTrigger
              trigger="click"
              placement="right"
              overlay={popover}
              rootClose
            >
              <Badge bg="secondary" className="p-1 fs-6 mx-2">
                Demo
              </Badge>
            </OverlayTrigger>
          </div>
          <AceEditor
            mode="yaml"
            theme="chrome"
            value={workflow}
            setOptions={{ useWorker: false }}
            onChange={setWorkflow}
            name="workflow-editor"
            editorProps={{ $blockScrolling: true }}
            width="153%"
            height="30em"
          />
        </div>
      </div>
      <Button
        className="float-end mt-3"
        variant="primary"
        onClick={handleSaveWorkflow}
      >
        Save
      </Button>
      <Button
        className="float-end mt-3 mx-1"
        variant="secondary"
        onClick={triggerFileInput}
      >
        Load File
      </Button>
      <input
        type="file"
        id="file-input"
        accept=".yaml,.yml"
        style={{ display: "none" }}
        onChange={handleLoadWorkflow}
      />
    </Container>
  );
}

export default WorkflowEditor;
