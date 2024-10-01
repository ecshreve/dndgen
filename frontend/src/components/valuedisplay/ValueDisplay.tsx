import React from "react";
import "./ValueDisplay.css";

const ValueDisplay = ({
  value,
  label,
  secondaryValue,
  highlight,
}: {
  value: number;
  label?: string;
  secondaryValue?: number;
  highlight?: boolean;
}) => {
  const displayBackground = highlight ? "yellow" : "";
  const circleBackgroundColor = value >= 0 ? "green" : "red";
  const squareBackgroundColor = "lightgray";

  return (
    <div className="value-display" style={{backgroundColor: displayBackground}}>
      <div
        style={{
          border: "1px solid black",
          padding: "10px",
          borderRadius: "10px",
        }}
      >
        <div className="label">{label}</div>
        <div
            style={{
              display: "flex",
            flexDirection: "row",
            alignItems: "center",
          }}
        >
        {secondaryValue && (
          <div
            className="square"
            style={{ backgroundColor: squareBackgroundColor }}
            >
              {secondaryValue}
            </div>
          )}
          <div
            className="circle"
            style={{ backgroundColor: circleBackgroundColor }}
            >
              {value >= 0 && "+"}
              {value}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ValueDisplay;
