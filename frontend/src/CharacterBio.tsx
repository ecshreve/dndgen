import { Box, Stack, TextField, Typography } from '@mui/material';
import React, { useState } from 'react';

type CharacterBioProps = {
  name: string;
  age: number;
  level: number;
  enableEdit: boolean;
  handleUpdate: (character: { name: string, age: number, level: number }) => void;
};

const CharacterBio = ({ name, age, level, enableEdit, handleUpdate }: CharacterBioProps) => {
  // Temporary state for editing values
  const [editValues, setEditValues] = useState({ name, age, level });

  // Handle input changes during editing
  const handleInputChange = (e: { target: { name: any; value: any; }; }) => {
    const { name, value } = e.target;
    setEditValues({ ...editValues, [name]: value })
    handleUpdate(editValues);
  };

  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Typography variant="h5" gutterBottom borderBottom="1px solid #ccc">
        Character Info
      </Typography>

      <Stack direction="row" spacing={2}>
        {/* Name Field */}
        <Box display="flex" alignItems="center" mb={2} flex={2} borderRadius={3} padding={1}>
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
            <Typography variant="body1"><strong>Name:</strong> {name}</Typography>
          )}
        </Box>

        {/* Age Field */}
        <Box display="flex" alignItems="center" mb={2} flex={1} borderRadius={3} padding={1}>
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
            <Typography variant="body1"><strong>Age:</strong> {age}</Typography>
          )}
          
        </Box>

        {/* Level Field */}
        <Box display="flex" alignItems="center" mb={2} flex={1} borderRadius={3} padding={1}>
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
            <Typography variant="body1"><strong>Level:</strong> {level}</Typography>
          )}
          
        </Box>
      </Stack>
    </Box>
  );
};

export default CharacterBio;
