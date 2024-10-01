import { Box } from "@mui/material";
import React from "react";
import "./SkillDisplay.css";

export type SkillDisplayProps = {
  skillName: string;
  abilityScoreName: string;
  abilityScoreModifier: number;
  proficient: boolean;
  proficiencyBonus: number;
}

const SkillDisplay = ({
  skillName,
  abilityScoreName,
  abilityScoreModifier,
  proficient,
  proficiencyBonus,
}: SkillDisplayProps) => {

  const asBackground = abilityScoreModifier > 0 ? "lightgreen" : abilityScoreModifier < 0 ? "lightcoral" : "lightgray";
  const fullModifier = abilityScoreModifier + (proficient ? proficiencyBonus : 0);
  const fullBackground = fullModifier > 0 ? "green" : fullModifier < 0 ? "red" : "lightgray";
  return (
    <div style={{display: "flex", flexDirection: "column"}}>
    <div className="skill-label-full">{skillName}</div>
    <Box className="skill-value-display">
      <Box sx={{display: "flex", flexDirection: "column", alignItems: "center"}}>
        <div className="as-circle" style={{backgroundColor: asBackground}}>
          {abilityScoreModifier > 0 && "+"}
          {abilityScoreModifier}
        </div>
        <div className="skill-label">{abilityScoreName}</div>
      </Box>
      {proficient && (
        <Box sx={{display: "flex", flexDirection: "column", alignItems: "center"}}>
          <div className="proficient-circle">+{proficiencyBonus}</div>
          <div className="skill-label">PROF</div>
        </Box>
      )}
      <Box>
        <div className="skill-circle" style={{backgroundColor: fullBackground}}>
          {fullModifier > 0 && "+"}
          {fullModifier}
        </div>
      </Box>
    </Box>
    </div>
  );
};

export default SkillDisplay;