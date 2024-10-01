import { useMutation, useQuery } from "@apollo/client";
import {
  Box,
  Button,
  CircularProgress,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Stack,
  TextField,
  Typography
} from "@mui/material";
import React, { useState } from "react";
import AbilityScoreDisplay from "./components/abilityscore/AbilityScoreDisplay";
import SkillDisplay from "./components/skilldisplay/SkillDisplay";
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
    id: string;
    indx: string;
    name: string;
  };
  characterAbilityScores: {
    id: string;
    score: number;
    modifier: number;
    abilityScore: {
      id: string;
      indx: string;
      name: string;
    };
  }[];
  characterSkills: {
    id: string;
    score: number;
    proficient: boolean;
    skill: {
      id: string;
      indx: string;
      name: string;
    };
    characterAbilityScore: {
      id: string;
      score: number;
      modifier: number;
      abilityScore: {
        id: string;
        indx: string;
        name: string;
      };
    };
  }[];
};

const CharacterSummary: React.FC = () => {
  const {
    loading: loadingCharacters,
    error: errorCharacters,
    data: characterData,
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

  const [
    updateCharacter,
    { data: updateData, loading: updateLoading, error: updateError },
  ] = useMutation(UPDATE_CHARACTER, {
    refetchQueries: [GET_CHARACTERS],
  });

  const [isEditing, setIsEditing] = useState(false);
  const [characterUpdate, setCharacterUpdate] =
    useState<CharacterDetails | null>(null);

  const [selectedAbilityScore, setSelectedAbilityScore] = useState<string | null>(null);
  const [selectedSkill, setSelectedSkill] = useState<string | null>(null);
  const [highlightedAbilityScore, setHighlightedAbilityScore] = useState<string | null>(null);
  const [highlightedSkills, setHighlightedSkills] = useState<string[]>([]);

  if (loadingCharacters || loadingClasses || loadingRaces || updateLoading)
    return <CircularProgress />;
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

  const character: CharacterDetails = characterData.characters[0];

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
            },
          },
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

  const handleSelectSkill = (skillName: string, abilityScoreName: string) => {
    setSelectedSkill(skillName);
    setHighlightedAbilityScore(abilityScoreName);
    setHighlightedSkills([]);
    setSelectedAbilityScore(null);
  };

  const handleSelectAbilityScore = (abilityScoreName: string) => {
    setSelectedAbilityScore(abilityScoreName);
    setHighlightedAbilityScore(null);
    setHighlightedSkills(character.characterSkills.filter(cs => cs.characterAbilityScore.abilityScore.name === abilityScoreName).map(cs => cs.skill.name));
    setSelectedSkill(null);
  };

  return (
    <Box p={2}>
      <Typography variant="h2">Character Summary</Typography>
      {character && (
        <>
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
          </Box>

          <Typography variant="h3">AbilityScores</Typography>
          <Box sx={{display: "flex", flexDirection: "row", justifyContent: "space-between"}}>
            {character.characterAbilityScores.map((cas) => (
              <Box key={cas.id}>
                <AbilityScoreDisplay
                  name={cas.abilityScore.name}
                  baseValue={cas.score}
                  modifier={cas.modifier}
                  highlighted={highlightedAbilityScore === cas.abilityScore.name}
                  selected={selectedAbilityScore === cas.abilityScore.name}
                  onSelect={handleSelectAbilityScore}
                />
              </Box>
            ))}
          </Box>

          <Typography variant="h3">Skills</Typography>
          <Box sx={{display: "flex", flexDirection: "row", justifyContent: "space-between", flexWrap: "wrap"}}>
            {character.characterSkills.map((cs) => (
              <SkillDisplay
                skillName={cs.skill.name}
                abilityScoreName={cs.characterAbilityScore.abilityScore.name}
                modifier={cs.characterAbilityScore.modifier}
                proficient={cs.proficient}
                bonus={character.proficiencyBonus}
                highlight={highlightedSkills.includes(cs.skill.name)}
                selected={selectedSkill === cs.skill.name}
                onSelect={() => handleSelectSkill(cs.skill.name, cs.characterAbilityScore.abilityScore.name)}
              />
            ))}
          </Box>
        </>
      )}
    </Box>
  );
};

export default CharacterSummary;
