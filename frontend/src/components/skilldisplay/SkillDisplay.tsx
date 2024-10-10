import { Box } from "@mui/material";
import React from "react";
import "./SkillDisplay.css";

export type SkillDisplayProps = {
  skillName: string;
  abilityScoreName: string;
  abilityScoreModifier: number;
  proficient: boolean;
  proficiencyBonus: number;
  shouldHighlight: boolean;
  selected: boolean;
  handleClick: () => void;
}

const SkillDisplay = ({
  skillName,
  abilityScoreName,
  abilityScoreModifier,
  proficient,
  proficiencyBonus,
  shouldHighlight,
  selected,
  handleClick
}: SkillDisplayProps) => {

  const componentBackground = selected ? "lightyellow" : shouldHighlight ? "yellow" : "lavender";
  const asBackground = abilityScoreModifier > 0 ? "lightgreen" : abilityScoreModifier < 0 ? "lightcoral" : "darkgray";
  const fullModifier = abilityScoreModifier + (proficient ? proficiencyBonus : 0);
  const fullBackground = fullModifier > 0 ? "green" : fullModifier < 0 ? "red" : "darkgray";
  return (
    <div className="skill-display" style={{backgroundColor: componentBackground}} onClick={handleClick}>
    <div className="skill-label-full">
      {skillName}
      {proficient && (
        <Box>
          <div className="proficient-circle">+{proficiencyBonus}</div>
        </Box>
      )}
    </div>
    <Box className="skill-value-display">
      <Box sx={{display: "flex", flexDirection: "column", alignItems: "center"}}>
        <div className="as-circle" style={{backgroundColor: asBackground}}>
          {abilityScoreModifier > 0 && "+"}
          {abilityScoreModifier}
        </div>
        <div className="skill-label">{abilityScoreName}</div>
      </Box>
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