import React from "react";
import "./SkillDisplay.css";

const SkillDisplay = ({
  skillName,
  abilityScoreName,
  modifier,
  proficient,
  bonus,
}: {
  skillName: string;
  abilityScoreName: string;
  modifier: number;
  proficient: boolean;
  bonus: number;
}) => {
  const asBackground = modifier > 0 ? "lightgreen" : "lightgray";
  const profBackground = "lightblue";
  const fullModifier = modifier + (proficient ? bonus : 0);
  const fullBackground =
    fullModifier > 0 ? "green" : fullModifier < 0 ? "red" : "lightgray";

  return (
    <div className="skill-value-display">
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          border: "1px solid black",
          borderRadius: "5px",
          padding: "5px",
        }}
      >
        <div className="skill-label">{skillName}</div>
        <div
          style={{
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            padding: "5px",
            borderTop: "1px solid black",
          }}
        >
          {proficient && (
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                padding: "5px",
              }}
            >
              <div className="proficient-circle">+{bonus}</div>
              <div className="skill-label">PROF</div>
            </div>
          )}
          <div
            style={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              padding: "5px",
            }}
          >
            <div
              className="small-circle"
              style={{ backgroundColor: asBackground }}
            >
              {modifier >= 0 ? "+" : ""}
              {modifier}
            </div>
            <div className="skill-label">{abilityScoreName}</div>
          </div>
          <div
            className="skill-circle"
            style={{ backgroundColor: fullBackground }}
          >
            {fullModifier >= 0 ? "+" : ""}
            {fullModifier}
          </div>
        </div>
      </div>
    </div>
  );
};

export default SkillDisplay;
