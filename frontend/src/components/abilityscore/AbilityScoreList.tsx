import { useQuery } from "@apollo/client";
import React, { useState } from "react";
import { CharacterAbilityScore } from "../../__generated__/graphql";
import { GET_ABILITY_SCORES } from "../../queries/getAbilityScores";
import AbilityScoreDisplay from "./AbilityScoreDisplay";

const AbilityScoreList = () => {
  const [selectedAbilityScore, setSelectedAbilityScore] = useState("")
  const { data, loading, error } = useQuery(GET_ABILITY_SCORES, {variables: {characterId: "12884901889"}});

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

  const handleAbilityScoreClick = (abilityScoreName: string) => {
    if (abilityScoreName === selectedAbilityScore) {
      setSelectedAbilityScore("");
    } else {
      setSelectedAbilityScore(abilityScoreName);
    }
  };
  
  return (
    <div style={{display: "flex", flexDirection: "row", justifyContent: "space-between"}}>
      {data && data.characters.edges[0].node.characterAbilityScores.map((cas: CharacterAbilityScore) => (
        <AbilityScoreDisplay
          key={cas.id}
          name={cas.abilityScore.name}
          baseValue={cas.score}
          modifier={cas.modifier}
          selected={selectedAbilityScore === cas.abilityScore.name}
          handleClick={() => {handleAbilityScoreClick(cas.abilityScore.name)}}
        />
      ))}
    </div>
  );
};

export default AbilityScoreList;
