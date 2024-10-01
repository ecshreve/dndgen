import React from "react";
import "./ValueDisplay.css";

const ValueDisplay = ({
  value,
  label,
  secondaryValue,
}: {
  value: number;
  label?: string;
  secondaryValue?: number;
}) => {
  const circleBackgroundColor = value >= 0 ? "green" : "red";
  const squareBackgroundColor = "lightgray";

  return (
    <div className="value-display">
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
