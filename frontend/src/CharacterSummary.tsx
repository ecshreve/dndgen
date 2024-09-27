import { useQuery } from "@apollo/client";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import RadioButtonUncheckedIcon from "@mui/icons-material/RadioButtonUnchecked";
import {
  Box,
  CircularProgress,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
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
  characterSkills: {
    modifier: number;
    proficient: boolean;
    skill: {
      indx: string;
      abilityScore: {
        indx: string;
      }
    }
  };
  characterAbilityScores: {
    modifier: number;
    score: number;
    abilityScore: {
      indx: string;
      skills: {
        characterSkills: {
          proficient: boolean;
          modifier: number;
          indx: string;
        }
      }
    };
  };
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
            maxWidth: "800px",
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
            <strong>Proficiency Bonus:</strong> +{character.proficiencyBonus}
          </Typography>
          <Typography variant="body1">
            <strong>Alignment:</strong> {character.alignment.name}
          </Typography>
          <Typography variant="body1">
            <strong>Ability Scores:</strong>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell>Ability</TableCell>
                  <TableCell>Base</TableCell>
                  <TableCell>Modifier</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {character.characterAbilityScores.map((abilityScore: any) => (
                  <TableRow key={abilityScore.abilityScore.indx}>
                    <TableCell>{abilityScore.abilityScore.indx}</TableCell>
                    <TableCell>{abilityScore.score}</TableCell>
                    <TableCell>
                      {abilityScore.modifier > 0 ? `+${abilityScore.modifier}` : abilityScore.modifier}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </Typography>
          <Typography variant="body1">
            <strong>Skills:</strong>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell>Skill</TableCell>
                  <TableCell>Proficient</TableCell>
                  <TableCell>Ability Score Modifier</TableCell>
                  <TableCell>Modifier</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {character.characterSkills.map((skill: any) => (
                  <TableRow key={skill.skill.indx}>
                    <TableCell>{skill.skill.indx}</TableCell>
                    <TableCell>
                      {skill.proficient ? (
                        <Stack
                          direction="row"
                          display="flex"
                          alignItems="center"
                        >
                          <CheckCircleIcon style={{ color: "green" }} />
                          <Typography variant="body2" marginLeft={1}>
                            +{character.proficiencyBonus}
                          </Typography>
                        </Stack>
                      ) : (
                        <RadioButtonUncheckedIcon style={{ color: "gray" }} />
                      )}
                    </TableCell>
                    <TableCell>
                      <Stack direction="row" spacing={1} alignItems="center">
                      <Typography>
                        {(() => {
                          const modifier = character.characterAbilityScores.find(
                            (abilityScore: any) =>
                              abilityScore.abilityScore.indx ===
                              skill.skill.abilityScore.indx
                          )?.modifier;
                          return modifier > 0 ? `+${modifier}` : modifier;
                        })()}
                      </Typography>
                        <Typography variant="body2" color="gray">
                        {
                          skill.skill.abilityScore.indx
                        }
                      </Typography>
                      </Stack>
                    </TableCell>
                    <TableCell>
                      {skill.modifier > 0 ? `+${skill.modifier}` : skill.modifier}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </Typography>
        </Box>
      )}
    </Box>
  );
};

export default CharacterSummary;
