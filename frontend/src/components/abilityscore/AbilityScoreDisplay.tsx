import React from "react";
import "./AbilityScoreDisplay.css";

type AbilityScoreDisplayProps = {
  name: string;
  baseValue: number;
  modifier: number;
  highlighted: boolean;
  selected: boolean;
  onSelect: (name: string) => void;
}



const AbilityScoreDisplay = (props: AbilityScoreDisplayProps) => {
  const circleBackgroundColor = props.modifier > 0 ? "lightgreen" : props.modifier < 0 ? "lightred" : "lightgray";
  const squareBackgroundColor = "lightgray";

  return (
    <div className={`ability-score-display ${props.selected ? "selected" : ""} ${props.highlighted ? "highlight" : ""}`} onClick={() => props.onSelect(props.name)}>
      <div className="label">{props.name}</div>
      <div className="ability-score-display-row">
        <div className="square" style={{backgroundColor: squareBackgroundColor}}>
                {props.baseValue}
            </div>
            <div className="circle" style={{backgroundColor: circleBackgroundColor}}>
                {props.modifier > 0 && "+"}
                {props.modifier}
            </div>
        </div>
    </div>
  );
};

export default AbilityScoreDisplay;