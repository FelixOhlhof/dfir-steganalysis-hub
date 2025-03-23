import "./App.css";
import WorkflowEditor from "./Components/WorkflowEditor";
import WorkflowDemo from "./Components/WorkflowDemo";
import { Container } from "react-bootstrap";
import { useAlert } from "react-alert";
import ServiceList from "./Components/ServiceList";
import React, { useEffect, useState } from "react";
import { transitions, positions, Provider as AlertProvider } from "react-alert";
import config from "./config";
import axios from "axios";

function App() {
  const [services, setServices] = useState([]);

  // const alert = useAlert();
  const servicesEndpoint = `${config.restGwUrl}/services`;

  useEffect(() => {
    axios
      .get(servicesEndpoint)
      .then((response) => {
        const formattedServices = response.data?.services?.map((service) => {
          return {
            ...service,
            functions: service.functions.map((func) => ({
              ...func,
              isExpanded: false,
            })),
          };
        });
        setServices(formattedServices);
        // console.log(response);
      })
      .catch((error) => {
        console.error("Error fetching services:", error);
        // alert.show(
        //   `Error fetching services: ${error} ${error?.response?.data}`,
        //   { type: "error" }
        // );
      });
  }, []);

  const options = {
    position: positions.BOTTOM_CENTER,
    timeout: 6000,
    offset: "30px",
    transition: transitions.SCALE,
  };

  const AlertTemplate = ({ message, options }) => {
    const alertClass =
      options.type === "success"
        ? "alert-success"
        : options.type === "error"
        ? "alert-danger"
        : "alert-warning";

    return (
      <div className={`alert ${alertClass} alert-dismissible mx-3`}>
        {message}
      </div>
    );
  };

  const style = {
    overflowY: "scroll",
  };

  return (
    <AlertProvider template={AlertTemplate} {...options}>
      <Container className="my-3">
        <div className="row">
          <div className="col-8">
            <WorkflowEditor />
            <WorkflowDemo />
          </div>
          <div className="col">
            <ServiceList services={services} setServices={setServices} />
          </div>
        </div>
      </Container>
    </AlertProvider>
  );
}

export default App;
