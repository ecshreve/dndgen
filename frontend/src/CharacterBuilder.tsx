import EditIcon from '@mui/icons-material/Edit';
import SaveIcon from '@mui/icons-material/Save';
import { Box, Button, Stack, Typography } from '@mui/material';
import React, { useState } from 'react';
import { ClassDetailsFragment, RaceDetailsFragment } from './__generated__/graphql';
import CharacterBio from './CharacterBio';
import AbilityScorePicker from './components/builder/abilityscorepicker';
import RacePicker from './components/builder/racepicker';
import SkillPicker from './components/builder/skillpicker';

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

  
const CharacterBuilderPage = () => {
  // Initial character data
  const [character, setCharacter] = useState({
    name: 'Zeke',
    age: 20,
    level: 5,
  });
  const [selectedClass, setSelectedClass] = useState<ClassDetailsFragment | null>(null);
  const [selectedRace, setSelectedRace] = useState<RaceDetailsFragment | null>(null);
  const [abilityScoreValues, setAbilityScoreValues] = useState<{ [key: string]: number }>({
    str: 8,
    dex: 8,
    con: 8,
    int: 8,
    wis: 8,
    cha: 8,
  });
  const [enableEdit, setEnableEdit] = useState(false);

  const handleSelectClass = (classDetails: ClassDetailsFragment) => {
    setSelectedClass(classDetails);
  };

  const handleSelectRace = (raceDetails: RaceDetailsFragment) => {
    setSelectedRace(raceDetails);
  };

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
      </Box>

      <Stack spacing={2}>
        <CharacterBio 
          name={character.name}
          age={character.age}
          level={character.level}
          enableEdit={enableEdit}
        />
        {/* <ClassPicker onSelect={handleSelectClass} /> */}
        <RacePicker onSelect={handleSelectRace}/>
        <AbilityScorePicker 
          scoreValues={abilityScoreValues}
          handleChange={(indx, value) => setAbilityScoreValues({ ...abilityScoreValues, [indx]: value })}
          enableEdit={enableEdit} />
        <SkillPicker 
          abilityScoreValues={abilityScoreValues}
          proficientSkills={selectedRace?.startingProficiencies?.map((prof: any) => prof.indx) || []}
          enableEdit={enableEdit} />
        
      </Stack>
    </Box>
  );
};

export default CharacterBuilderPage;