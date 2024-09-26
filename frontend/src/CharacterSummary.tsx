import { useQuery } from "@apollo/client";
import { Box, CircularProgress, Typography } from "@mui/material";
import React from "react";
import { GET_CHARACTERS } from "./queries/getCharacters";

type CharacterDetails = {
  id: string;
  name: string;
  age: number;
  level: number;
  race: {
    name: string;
  };
  class: {
    name: string;
  };
  proficiencyBonus: number;
  alignment: {
    name: string;
  };
  characterSkills: any;
  characterAbilityScores: any;
};

const CharacterSummary: React.FC = () => {
  const { loading, error, data } = useQuery(GET_CHARACTERS);

  if (loading) return <CircularProgress />;
  if (error) return <p>SUMMARYError: {error.message}</p>;

  const characters = data.characters.edges.map(
    (edge: { node: CharacterDetails }) => edge.node
  );

  const character = characters[0];

  return (
    <Box p={2}>
      {character && (
        <Box
          sx={{
            maxWidth: "500px",
            margin: "0 auto",
            padding: "20px",
            border: "1px solid #ccc",
            borderRadius: "10px",
            backgroundColor: "#f9f9f9",
          }}
        >
          <Typography variant="h4" gutterBottom>
            {character.name}
          </Typography>
          <Typography variant="body1">
            <strong>Class:</strong> {character.class.name}
          </Typography>
          <Typography variant="body1">
            <strong>Race:</strong> {character.race.name}
          </Typography>
          <Typography variant="body1">
            <strong>Level:</strong> {character.level}
          </Typography>
          <Typography variant="body1">
            <strong>Age:</strong> {character.age}
          </Typography>
          <Typography variant="body1">
            <strong>Proficiency Bonus:</strong> {character.proficiencyBonus}
          </Typography>
          <Typography variant="body1">
            <strong>Alignment:</strong> {character.alignment.name}
          </Typography>
          <Typography variant="body1">
            <strong>Ability Scores:</strong>
            <ul>
              {character.characterAbilityScores.map((abilityScore: any) => (
                <li key={abilityScore.abilityScore.indx}>
                  {abilityScore.abilityScore.indx} - Base: {abilityScore.score} - Modifier: {abilityScore.modifier}
                </li>
              ))}
            </ul>
          </Typography>
          <Typography variant="body1">
            <strong>Skills:</strong>
            <ul>
              {character.characterSkills.map((skill: any) => (
                <li key={skill.skill.indx}>
                  {skill.skill.indx} - Proficient:{" "}
                  {skill.proficient ? "Yes" : "No"}, Modifier: {skill.modifier}
                </li>
              ))}
            </ul>
          </Typography>
        </Box>
      )}
    </Box>
  );
};

export default CharacterSummary;
