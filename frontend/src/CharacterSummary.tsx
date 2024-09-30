import { useMutation, useQuery } from "@apollo/client";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import RadioButtonUncheckedIcon from "@mui/icons-material/RadioButtonUnchecked";
import {
  Box,
  Button,
  CircularProgress,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from "@mui/material";
import React, { useState } from "react";
import ValueDisplay from "./components/valuedisplay/ValueDisplay";
import { GET_CHARACTERS } from "./queries/getCharacters";
import { GET_CLASSES } from "./queries/getClasses";
import { GET_RACES } from "./queries/getRaces";
import { UPDATE_CHARACTER } from "./queries/updateCharacter";

interface ClassDetails {
  id: string;
  indx: string;
  name: string;
}

interface RaceDetails {
  id: string;
  indx: string;
  name: string;
}

type CharacterDetails = {
  id: string;
  name: string;
  age: number;
  level: number;
  race: {
    id: string;
    indx: string;
    name: string;
  };
  class: {
    id: string;
    indx: string;
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
      };
    };
  }[];
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
        };
      };
    };
  }[];
};

const CharacterSummary: React.FC = () => {
  const { 
    loading: loadingCharacters, 
    error: errorCharacters, 
    data: characterData
  } = useQuery(GET_CHARACTERS);

  // Fetching classes
  const {
    loading: loadingClasses,
    error: errorClasses,
    data: classData,
  } = useQuery(GET_CLASSES);

  // Fetching races
  const {
    loading: loadingRaces,
    error: errorRaces,
    data: raceData,
  } = useQuery(GET_RACES);

  const [updateCharacter, {data: updateData, loading: updateLoading, error: updateError }] = useMutation(UPDATE_CHARACTER,
    {
      refetchQueries: [GET_CHARACTERS],
    }
  );

  const [isEditing, setIsEditing] = useState(false);
  const [characterUpdate, setCharacterUpdate] =
    useState<CharacterDetails | null>(null);

  if (loadingCharacters || loadingClasses || loadingRaces || updateLoading) return <CircularProgress />;
  if (errorCharacters || errorClasses || errorRaces || updateError)
    return (
      <>
        <p>Error loading data...</p>
        <p>{errorCharacters?.message}</p>
        <p>{errorClasses?.message}</p>
        <p>{errorRaces?.message}</p>
        <p>{updateError?.message}</p>
      </>
    );

  const character: CharacterDetails = characterData.characters.edges.map(
    (edge: { node: CharacterDetails }) => edge.node
  )[0];

  const handleEditClick = async () => {
    if (isEditing) {
      if (characterUpdate) {
        const updated = await updateCharacter({
          variables: {
            updateCharacterId: characterUpdate.id,
            input: {
              name: characterUpdate.name,
              age: characterUpdate.age,
              level: characterUpdate.level,
              raceID: characterUpdate.race.id,
              classID: characterUpdate.class.id,
            }
          }
        });
        console.log(updated); 
      }
    }
    setIsEditing(!isEditing);
    setCharacterUpdate(character);
  };

  const handleChange = (e: any) => {
    if (characterUpdate) {
      if (e.target.name === "race") {
        setCharacterUpdate({
          ...characterUpdate,
          race: raceData.races.find(
            (race: RaceDetails) => race.indx === e.target.value
          ),
        });
      } else if (e.target.name === "class") {
          setCharacterUpdate({
          ...characterUpdate,
          class: classData.classes.find(
            (cls: ClassDetails) => cls.indx === e.target.value
          ),
        });
      } else {
        setCharacterUpdate({
          ...characterUpdate,
          [e.target.name]: e.target.value,
        });
      }
    }
  };

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
          <Stack
            direction="row"
            justifyContent="space-between"
            alignItems="center"
          >
            <Typography variant="h4" gutterBottom>
              {isEditing ? (
                <TextField
                  variant="outlined"
                  name="name"
                  value={characterUpdate?.name || ""}
                  onChange={handleChange}
                  fullWidth
                />
              ) : (
                character.name
              )}
            </Typography>
            <Button onClick={handleEditClick}>
              {isEditing ? "Save" : "Edit"}
            </Button>
          </Stack>
          <Typography variant="body1">
            <strong>Class:</strong>{" "}
            {isEditing ? (
              <FormControl fullWidth margin="normal">
                <InputLabel id="class-select-label">Select Class</InputLabel>
                <Select
                  labelId="class-select-label"
                  value={characterUpdate?.class.indx || ""}
                  onChange={handleChange}
                  label="Select a class"
                  name="class"
                >
                  {classData.classes.map((cls: ClassDetails) => (
                    <MenuItem key={cls.id} value={cls.indx}>
                      {cls.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            ) : (
              character.class.name
            )}
          </Typography>

          <Typography variant="body1">
            <strong>Race:</strong>{" "}
            {isEditing ? (
              <FormControl fullWidth margin="normal">
                <InputLabel id="race-select-label">Select Race</InputLabel>
                <Select
                  labelId="race-select-label"
                  value={characterUpdate?.race.indx || ""}
                  onChange={handleChange}
                  name="race"
                  label="Select a race"
                >
                  {raceData.races.map((race: RaceDetails) => (
                    <MenuItem key={race.id} value={race.indx}>
                      {race.name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            ) : (
              character.race.name
            )}
          </Typography>
          <Typography variant="body1">
            <strong>Level:</strong>{" "}
            {isEditing ? (
              <TextField
                variant="outlined"
                type="number"
                name="level"
                value={characterUpdate?.level || ""}
                onChange={handleChange}
                fullWidth
              />
            ) : (
              character.level
            )}
          </Typography>
          <Typography variant="body1">
            <strong>Age:</strong>{" "}
            {isEditing ? (
              <TextField
                variant="outlined"
                type="number"
                name="age"
                value={characterUpdate?.age || ""}
                onChange={handleChange}
                fullWidth
              />
            ) : (
              character.age
            )}
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
                {character.characterAbilityScores.map((cas: any) => (
                  <TableRow key={cas.abilityScore.indx}>
                    <TableCell>{cas.abilityScore.indx}</TableCell>
                    <TableCell>{cas.score}</TableCell>
                    <TableCell>
                      <ValueDisplay value={cas.modifier} label={cas.abilityScore.indx} />
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
                {character.characterSkills.map((cas: any) => (
                  <TableRow key={cas.skill.indx}>
                    <TableCell>
                      {cas.skill.indx}
                    </TableCell>
                    <TableCell>
                      {cas.proficient ? (
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
                      <ValueDisplay value={character.characterAbilityScores.find(
                        (caa: any) => caa.abilityScore.indx === cas.skill.abilityScore.indx
                      )?.modifier || 0} label={cas.skill.abilityScore.indx} />
                    </TableCell>
                    <TableCell>
                      <ValueDisplay value={character.characterSkills.find(
                        (caa: any) => caa.skill.indx === cas.skill.indx
                      )?.modifier || 0} />
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
