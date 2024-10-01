import { Box, Button, Typography } from '@mui/material';
import React, { useState } from 'react';
import CharacterBio from './CharacterBio';
import AbilityScorePicker from './components/builder/abilityscorepicker';
import ClassPicker from './components/builder/classpicker';
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
  
  interface ClassDetails {
    id: string;
    indx: string;
    name: string;
    savingThrows: { indx: string }[];
    proficiencies: { indx: string }[];
    proficiencyOptions: { desc: string }[];
    startingEquipment: { quantity: number; equipment: { indx: string } }[];
  }
  
const CharacterBuilderPage = () => {
  // Initial character data
  const [character, setCharacter] = useState({
    name: 'Zeke',
    age: 20,
    level: 5,
    description: 'Zeke is a wildcard.',
  });

  const [enableEdit, setEnableEdit] = useState(false);

  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Box display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h4" gutterBottom>
          Character Builder
        </Typography>
        <Button variant="contained" color="primary" onClick={() => setEnableEdit(!enableEdit)}>
          {enableEdit ? 'Disable Edit' : 'Enable Edit'}
        </Button>
      </Box>

      <CharacterBio 
        name={character.name}
        age={character.age}
        level={character.level}
        description={character.description}
        enableEdit={enableEdit}
      />
      <AbilityScorePicker />
      <ClassPicker />
    </Box>
  );
};

export default CharacterBuilderPage;
