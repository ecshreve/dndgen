import { Box, Stack, TextField, Typography } from '@mui/material';
import React, { useState } from 'react';

type CharacterBioProps = {
  name: string;
  age: number;
  level: number;
  enableEdit: boolean;
};

const CharacterBio = ({ name, age, level, enableEdit }: CharacterBioProps) => {
  const [character, setCharacter] = useState({
    name: name,
    age: age,
    level: level,
  });
  
  // State to track which field is being edited
  const [editField, setEditField] = useState({
    name: false,
    age: false,
    level: false,
    description: false,
  });

  // Temporary state for editing values
  const [editValues, setEditValues] = useState({ ...character });

  // Handle input changes during editing
  const handleInputChange = (e: { target: { name: any; value: any; }; }) => {
    const { name, value } = e.target;
    setEditValues({ ...editValues, [name]: value });
  };

  // Toggle edit mode for a specific field
  const toggleEditField = (field: keyof typeof editField) => {
    setEditField({ ...editField, [field]: !editField[field] });

    // If we're leaving edit mode, save the changes
    if (editField[field]) {
      setCharacter(editValues);
    }
  };

  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Typography variant="h5" gutterBottom borderBottom="1px solid #ccc">
        Character Info
      </Typography>

      <Stack direction="row" spacing={2}>
        {/* Name Field */}
        <Box display="flex" alignItems="center" mb={2} flex={2} border="1px solid #ccc" borderRadius={3} padding={1}>
          {enableEdit ? (
            <TextField
              fullWidth
              label="Name"
              name="name"
              value={editValues.name}
              onChange={handleInputChange}
              margin="normal"
            />
          ) : (
            <Typography variant="body1"><strong>Name:</strong> {character.name}</Typography>
          )}
        </Box>

        {/* Age Field */}
        <Box display="flex" alignItems="center" mb={2} flex={1} border="1px solid #ccc" borderRadius={3} padding={1}>
          {enableEdit ? (
            <TextField
              fullWidth
              label="Age"
              name="age"
              value={editValues.age}
              onChange={handleInputChange}
              margin="normal"
              type="number"
            />
          ) : (
            <Typography variant="body1"><strong>Age:</strong> {character.age}</Typography>
          )}
          
        </Box>

        {/* Level Field */}
        <Box display="flex" alignItems="center" mb={2} flex={1} border="1px solid #ccc" borderRadius={3} padding={1}>
          {enableEdit ? (
            <TextField
              fullWidth
              label="Level"
              type="number"
              name="level"
              value={editValues.level}
              onChange={handleInputChange}
              margin="normal"
            />
          ) : (
            <Typography variant="body1"><strong>Level:</strong> {character.level}</Typography>
          )}
          
        </Box>
      </Stack>
    </Box>
  );
};

export default CharacterBio;
