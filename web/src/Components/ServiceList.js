import React, { useState } from "react";
import { Container, Tooltip, Collapse, OverlayTrigger } from "react-bootstrap";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faChevronDown,
  faChevronRight,
} from "@fortawesome/free-solid-svg-icons";

const types = {
  0: "string",
  1: "int",
  2: "float",
  3: "bool",
  4: "dict",
  5: "list",
  6: "bytes",
};

export default function ServiceList(props) {
  const [expandedParameterIndex, setExpandedParameterIndex] = useState(null);
  const [showIndex, setShowIndex] = useState({});

  const handleMouseEnter = (key) => {
    setShowIndex((prev) => ({ ...prev, [key]: true }));
  };

  const handleMouseLeave = (key) => {
    setShowIndex((prev) => ({ ...prev, [key]: false }));
  };
  const toggleExtendFunc = (serviceName, funcName) => {
    const updatedServices = props.services.map((service) => {
      if (service.name === serviceName) {
        const updatedFunctions = service.functions.map((func) => {
          if (func.name === funcName) {
            return {
              ...func,
              isExpanded: !func.isExpanded,
            };
          }
          return func;
        });
        return {
          ...service,
          functions: updatedFunctions,
        };
      }
      return service;
    });

    props.setServices(updatedServices);
  };

  return (
    <div>
      <Container className="pt-0">
        <h2 className="mb-4">Services</h2>
      </Container>
      <div style={{ overflowY: "auto", maxHeight: "80vh" }}>
        {props.services?.map((s, i) => (
          <Container
            key={i}
            className="p-3 mb-4 border rounded bg-body-secondary"
          >
            <h5 className="alert-heading d-flex align-items-center">
              <OverlayTrigger
                placement="top"
                show={showIndex[`service-${i}`]}
                overlay={
                  s.description && (
                    <Tooltip
                      id={`tooltip-${i}`}
                      onMouseEnter={() => handleMouseEnter(`service-${i}`)}
                      onMouseLeave={() => handleMouseLeave(`service-${i}`)}
                    >
                      {s.description}
                    </Tooltip>
                  )
                }
              >
                <span
                  onMouseEnter={() => handleMouseEnter(`service-${i}`)}
                  onMouseLeave={() => handleMouseLeave(`service-${i}`)}
                >
                  {s.name || "Unknown Service"}
                </span>
              </OverlayTrigger>
            </h5>
            <hr className="my-2" />

            {/* Functions */}
            {s.functions.map((f, fi) => (
              <div key={fi} className="mb-3">
                {/* Function Header */}
                <div className="row">
                  <div className="col-5 d-flex align-items-center">
                    <FontAwesomeIcon
                      icon={f.isExpanded ? faChevronDown : faChevronRight}
                      className="me-2"
                      onClick={() => toggleExtendFunc(s.name, f.name)}
                    />
                    <strong>{f.name}:</strong>
                  </div>
                  <div className="col-7">{f.description}</div>
                </div>

                {/* Expanded Content */}
                <Collapse in={f.isExpanded}>
                  <div className="mt-2">
                    {/* Return Fields */}
                    {f.return_fields && (
                      <div className="mb-3">
                        <h6>Return Fields:</h6>
                        <ul className="list-group list-group-flush">
                          {f.return_fields.map((rf, rfi) => (
                            <OverlayTrigger
                              key={rfi}
                              placement="top"
                              show={showIndex[`return-${i}-${fi}-${rfi}`]}
                              overlay={
                                rf.description && (
                                  <Tooltip
                                    id={`tooltip-rf-${rfi}`}
                                    onMouseEnter={() =>
                                      handleMouseEnter(
                                        `return-${i}-${fi}-${rfi}`
                                      )
                                    }
                                    onMouseLeave={() =>
                                      handleMouseLeave(
                                        `return-${i}-${fi}-${rfi}`
                                      )
                                    }
                                  >
                                    {rf.description}
                                  </Tooltip>
                                )
                              }
                            >
                              <li
                                className="list-group-item"
                                onMouseEnter={() =>
                                  handleMouseEnter(`return-${i}-${fi}-${rfi}`)
                                }
                                onMouseLeave={() =>
                                  handleMouseLeave(`return-${i}-${fi}-${rfi}`)
                                }
                              >
                                {rf.name}{" "}
                                {rf.type ? (
                                  <span className="text-muted">
                                    [{types[rf.type]}]
                                  </span>
                                ) : (
                                  <span className="text-muted">[string]</span>
                                )}
                              </li>
                            </OverlayTrigger>
                          ))}
                        </ul>
                      </div>
                    )}

                    {/* Parameters */}
                    {f.parameter && (
                      <div className="mb-3">
                        <h6>Parameters:</h6>
                        <ul className="list-group list-group-flush">
                          {f.parameter.map((p, pi) => (
                            <OverlayTrigger
                              key={pi}
                              placement="top"
                              show={showIndex[`param-${i}-${fi}-${pi}`]}
                              overlay={
                                p.description && (
                                  <Tooltip
                                    id={`tooltip-param-${pi}`}
                                    onMouseEnter={() =>
                                      handleMouseEnter(`param-${i}-${fi}-${pi}`)
                                    }
                                    onMouseLeave={() =>
                                      handleMouseLeave(`param-${i}-${fi}-${pi}`)
                                    }
                                  >
                                    {p.description
                                      .split("\n")
                                      .map((line, index) => (
                                        <span key={index}>
                                          {line}
                                          {index <
                                            p.description.split("\n").length -
                                              1 && <br />}
                                        </span>
                                      ))}
                                    {p.default && (
                                      <div>
                                        <strong>Default:</strong> {p.default}
                                      </div>
                                    )}
                                  </Tooltip>
                                )
                              }
                            >
                              <li
                                className="list-group-item"
                                onMouseEnter={() =>
                                  handleMouseEnter(`param-${i}-${fi}-${pi}`)
                                }
                                onMouseLeave={() =>
                                  handleMouseLeave(`param-${i}-${fi}-${pi}`)
                                }
                              >
                                {p.name}{" "}
                                {p.type ? (
                                  <span className="text-muted">
                                    [{types[p.type]}]
                                  </span>
                                ) : (
                                  <span className="text-muted">[string]</span>
                                )}{" "}
                                {p.optional && (
                                  <span className="text-muted">(optional)</span>
                                )}
                              </li>
                            </OverlayTrigger>
                          ))}
                        </ul>
                      </div>
                    )}

                    {/* Supported File Types */}
                    {f.supported_file_types && (
                      <div>
                        <h6>Supported File Types:</h6>
                        <ul className="list-group list-group-flush">
                          {f.supported_file_types.map((ft, fti) => (
                            <li key={fti} className="list-group-item">
                              {ft}
                            </li>
                          ))}
                        </ul>
                      </div>
                    )}
                  </div>
                </Collapse>
              </div>
            ))}
          </Container>
        ))}
      </div>
    </div>
  );
}
