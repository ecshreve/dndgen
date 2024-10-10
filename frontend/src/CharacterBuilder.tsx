import { useMutation } from '@apollo/client';
import EditIcon from '@mui/icons-material/Edit';
import SaveIcon from '@mui/icons-material/Save';
import StarIcon from '@mui/icons-material/Star';
import { Box, Button, Stack, Typography } from '@mui/material';
import React, { useState } from 'react';
import { ClassDetailsFragment, RaceDetailsFragment } from './__generated__/graphql';
import CharacterBio from './CharacterBio';
import AbilityScorePicker from './components/builder/abilityscorepicker';
import ClassPicker from './components/builder/classpicker';
import RacePicker from './components/builder/racepicker';
import SkillPicker from './components/builder/skillpicker';
import { CREATE_CHARACTER } from './queries/createCharacter';
import { calculateProfBonus } from './utils';

interface RaceDetails {
    id: string;
    indx: string;
    name: string;
    abilityBonuses: { bonus: number; abilityScore: { indx: string } }[];
    alignmentDesc: string;
    ageDesc: string;
    size: string;
    sizeDesc: string;
    speed: number;
    languages: { name: string }[];
    languageDesc: string;
    startingProficiencies: { indx: string }[];
    startingProficiencyOptions: {
      desc: string[];
      choose: number;
      proficiencies: { indx: string }[];
    };
  }

interface CharacterDetails {
  name: string;
  age: number;
  level: number;
  abilityScores: {
    str: number;
    dex: number;
    con: number;
    int: number;
    wis: number;
    cha: number;
  };
  race: {
    id: string;
    indx: string;
    name: string;
  } | null;
  class: {
    id: string;
    indx: string;
    name: string;
  } | null;
  proficiencies: string[];
}
const initialCharacter: CharacterDetails = {
  name: 'Zeke',
  age: 20,
  level: 5,
  abilityScores: {
    str: 8,
    dex: 8,
    con: 8,
    int: 8,
    wis: 8,
    cha: 8,
  },
  race: null,
  class: null,
  proficiencies: [],
};
  
const CharacterBuilderPage = () => {
  const [character, setCharacter] = useState<CharacterDetails>(initialCharacter);
  const [selectedClass, setSelectedClass] = useState<ClassDetailsFragment | null>(null);
  const [selectedRace, setSelectedRace] = useState<RaceDetailsFragment | null>(null);
  const [abilityScoreValues, setAbilityScoreValues] = useState<{ [key: string]: number }>(initialCharacter.abilityScores);
  const [selectedProficiencyOptions, setSelectedProficiencyOptions] = useState<{ race: string[], class: string[] }>({ race: [], class: [] });
  const [enableEdit, setEnableEdit] = useState(false);
  
  const [saveCharacter, {error, data}] = 
    useMutation(CREATE_CHARACTER, {
      variables: {
        input: {  
          name: character.name,
          age: character.age,
          level: character.level,
          raceID: character.race?.id,
          classID: character.class?.id,
        }
      }
    });

  

  const combinedProficiencies = [
    ...(selectedRace?.startingProficiencies?.map((prof: any) => prof.indx) || []),
    ...(selectedClass?.proficiencies?.map((prof: any) => prof.indx) || []),
    ...selectedProficiencyOptions.race,
    ...selectedProficiencyOptions.class,
  ];

  const handleSelectClass = (classDetails: ClassDetailsFragment) => {
    setSelectedProficiencyOptions({...selectedProficiencyOptions, class: []});
    setSelectedClass(classDetails);
    setCharacter({ ...character, class: classDetails });
  }

  const handleSelectRace = (raceDetails: RaceDetailsFragment) => {
    setSelectedProficiencyOptions({...selectedProficiencyOptions, race: []});
    setSelectedRace(raceDetails);
    setCharacter({ ...character, race: raceDetails });
  }

  const handleSelectProficiencyOptions = (options: string[], type: 'race' | 'class') => {
    setSelectedProficiencyOptions({ ...selectedProficiencyOptions, [type]: options });
    setCharacter({ ...character, proficiencies: combinedProficiencies });
  }

  const handleCharacterBioUpdate = (ch: { name: string, age: number, level: number }) => {
    setCharacter({ ...character, ...ch });
  }

  const handleCreateCharacter = () => {
    console.log(character);
    saveCharacter();
  }

  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h4" gutterBottom>
          Character Builder
        </Typography>
        <Button variant="contained" color="primary" onClick={() => setEnableEdit(!enableEdit)}>
          {enableEdit ? <SaveIcon /> : <EditIcon />}
          <Box component="span" marginLeft={1}>{enableEdit ? 'Save' : 'Edit'}</Box>
        </Button>
        <Button variant="contained" color="primary" onClick={handleCreateCharacter}>
          <StarIcon />
          <Box component="span" marginLeft={1}>Create Character</Box>
        </Button>
      </Box>

      <Stack spacing={2}>
        <CharacterBio 
          name={character.name}
          age={character.age}
          level={character.level}
          enableEdit={enableEdit}
          handleUpdate={handleCharacterBioUpdate}
        />
        <AbilityScorePicker 
          scoreValues={abilityScoreValues}
          handleChange={(indx, value) => setAbilityScoreValues({ ...abilityScoreValues, [indx]: value })}
          enableEdit={enableEdit} />
        <Box display="flex" flexDirection="row" gap={1}>
          <Box display="flex" flexDirection="column" flex={1}>
            <RacePicker 
              onSelect={handleSelectRace} 
              onProficiencyOptionsChange={(options) => handleSelectProficiencyOptions(options, 'race')}
              enableEdit={enableEdit}
            />
          </Box>
          <Box display="flex" flexDirection="column" flex={1}>
            <ClassPicker 
              onSelect={handleSelectClass} 
              onProficiencyOptionsChange={(options) => handleSelectProficiencyOptions(options, 'class')}
              enableEdit={enableEdit}
            />
          </Box>
        </Box>
        <SkillPicker 
          abilityScoreValues={abilityScoreValues}
          proficiencies={combinedProficiencies}
          profBonus={calculateProfBonus(character.level)}
          enableEdit={enableEdit} />
      </Stack>
    </Box>
  );
};

export default CharacterBuilderPage;
