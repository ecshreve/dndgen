import { Box } from "@mui/material";
import React from "react";
import "./AbilityScoreDisplay.css";

type AbilityScoreDisplayProps = {
  name: string;
  baseValue: number;
  modifier: number;
  selected: boolean;
  handleClick: () => void;
  shouldHighlight: boolean;
}



const AbilityScoreDisplay = (props: AbilityScoreDisplayProps) => {
  const backgroundColor = props.selected ? "lightyellow" : props.shouldHighlight ? "yellow" : "white";
  const circleBackgroundColor = props.modifier > 0 ? "lightgreen" : props.modifier < 0 ? "lightcoral" : "lightgray";
  const squareBackgroundColor = "lightgray";

  return (
    <div className={`ability-score-display`} style={{backgroundColor: backgroundColor}} onClick={props.handleClick}>
        <div className="ability-score-display-row">
            <Box sx={{display: "flex", flexDirection: "column", alignItems: "center", gap: "5px"}}>
                <div className="ability-score-label">{props.name}</div>
                <div className="ability-score-square" style={{backgroundColor: squareBackgroundColor}}>
                    {props.baseValue}
                </div>
            </Box>
            <div className="ability-score-circle" style={{backgroundColor: circleBackgroundColor}}>
                {props.modifier > 0 && "+"}
                {props.modifier}
            </div>
        </div>
    </div>
  );
};

export default AbilityScoreDisplay;