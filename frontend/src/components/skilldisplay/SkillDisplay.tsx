import React from "react";
import "./SkillDisplay.css";

const SkillDisplay = ({
  skillName,
  abilityScoreName,
  modifier,
  proficient,
  bonus,
  highlight,
  selected,
  onSelect
}: {
  skillName: string;
  abilityScoreName: string;
  modifier: number;
  proficient: boolean;
  bonus: number;
  highlight: boolean;
  selected: boolean;
  onSelect: () => void;
}) => {
  const asBackground = modifier > 0 ? "lightgreen" : "lightred";
  const fullModifier = modifier + (proficient ? bonus : 0);
  const fullModBackground =
    fullModifier > 0 ? "green" : fullModifier < 0 ? "red" : "lightgray";

  return (
    <div className={`skill-value-display ${highlight ? "highlight" : ""} ${selected ? "selected" : ""}`} onClick={onSelect}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
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
            style={{ backgroundColor: fullModBackground }}
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
