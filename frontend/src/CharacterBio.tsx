import EditIcon from '@mui/icons-material/Edit';
import SaveIcon from '@mui/icons-material/Save';
import { Box, IconButton, TextField, Typography } from '@mui/material';
import React, { useState } from 'react';

const CharacterBioPage = () => {
  // Initial character data
  const [character, setCharacter] = useState({
    name: 'Zeke',
    classType: 'Rogue',
    level: 5,
    description: 'Zeke is a wildcard.',
  });

  // State to track which field is being edited
  const [editField, setEditField] = useState({
    name: false,
    classType: false,
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
    <Box sx={{ maxWidth: '800px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Typography variant="h4" gutterBottom>
        Character Bio
      </Typography>

      {/* Name Field */}
      <Box display="flex" alignItems="center" mb={2} justifyContent="space-between" border="1px solid #ccc" borderRadius={3} padding={1}>
        {editField.name ? (
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
        <IconButton onClick={() => toggleEditField('name')}>
          {editField.name ? <SaveIcon /> : <EditIcon />}
        </IconButton>
      </Box>

      {/* Class Field */}
      <Box display="flex" alignItems="center" mb={2} justifyContent="space-between" border="1px solid #ccc" borderRadius={3} padding={1}>
        {editField.classType ? (
          <TextField
            fullWidth
            label="Class"
            name="classType"
            value={editValues.classType}
            onChange={handleInputChange}
            margin="normal"
          />
        ) : (
          <Typography variant="body1"><strong>Class:</strong> {character.classType}</Typography>
        )}
        <IconButton onClick={() => toggleEditField('classType')}>
          {editField.classType ? <SaveIcon /> : <EditIcon />}
        </IconButton>
      </Box>

      {/* Level Field */}
      <Box display="flex" alignItems="center" mb={2} justifyContent="space-between" border="1px solid #ccc" borderRadius={3} padding={1}>
        {editField.level ? (
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
        <IconButton onClick={() => toggleEditField('level')}>
          {editField.level ? <SaveIcon /> : <EditIcon />}
        </IconButton>
      </Box>

      {/* Description Field */}
      <Box display="flex" alignItems="center" mb={2} justifyContent="space-between" border="1px solid #ccc" borderRadius={3} padding={1}>
        {editField.description ? (
          <TextField
            fullWidth
            label="Description"
            name="description"
            value={editValues.description}
            onChange={handleInputChange}
            multiline
            rows={4}
            margin="normal"
          />
        ) : (
          <Typography variant="body1"><strong>Description:</strong> {character.description}</Typography>
        )}
        <IconButton onClick={() => toggleEditField('description')}>
          {editField.description ? <SaveIcon /> : <EditIcon />}
        </IconButton>
      </Box>
    </Box>
  );
};

export default CharacterBioPage;
