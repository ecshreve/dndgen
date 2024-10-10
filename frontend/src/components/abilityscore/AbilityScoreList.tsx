import { useQuery } from "@apollo/client";
import React from "react";
import { CharacterAbilityScore } from "../../__generated__/graphql";
import { GET_ABILITY_SCORES } from "../../queries/getAbilityScores";
import AbilityScoreDisplay from "./AbilityScoreDisplay";

interface AbilityScoreListProps {
  selectedAbilityScore: string;
  selectedSkill: string;
  onSelect: (abilityScore: string) => void;
}

const AbilityScoreList = ({ selectedAbilityScore, selectedSkill, onSelect }: AbilityScoreListProps) => {
  const { data, loading, error } = useQuery(GET_ABILITY_SCORES, {variables: {characterId: "12884901889"}});

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;
  
  return (
    <div style={{display: "flex", flexDirection: "row", justifyContent: "space-between"}}>
      {data && data.characters.edges[0].node.characterAbilityScores.map((cas: CharacterAbilityScore) => (
        <AbilityScoreDisplay
          key={cas.id}
          name={cas.abilityScore.name}
          baseValue={cas.score}
          modifier={cas.modifier}
          selected={selectedAbilityScore === cas.abilityScore.name}
          handleClick={() => {onSelect(cas.abilityScore.name)}}
          shouldHighlight={cas.characterSkills?.find(cs => cs.skill.name === selectedSkill) !== undefined}
        />
      ))}
    </div>
  );
};

export default AbilityScoreList;
