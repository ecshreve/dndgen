import { useQuery } from "@apollo/client";
import { Box } from "@mui/material";
import React from "react";
import { CharacterSkill } from "../../__generated__/graphql";
import { GET_SKILLS } from "../../queries/getSkills";
import SkillDisplay from "./SkillDisplay";
import "./SkillDisplay.css";


const SkillList = () => {
  const { data, loading, error } = useQuery(GET_SKILLS, {variables: {characterId: "12884901889"}});

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;


  return (
    <Box sx={{display: "flex", flexDirection: "row", gap: 2, flexWrap: "wrap"}}>
      {data && data.characters.edges[0].node.characterSkills.map((skill: CharacterSkill) => (
        <SkillDisplay
          key={skill.id}
          skillName={skill.skill.name}
          abilityScoreName={skill.characterAbilityScore!.abilityScore.name}
          abilityScoreModifier={skill.characterAbilityScore!.modifier}
          proficient={skill.proficient}
          proficiencyBonus={data.characters.edges[0].node.proficiencyBonus}
        />
      ))}
    </Box>
  );
};

export default SkillList;
