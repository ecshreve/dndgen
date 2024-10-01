import EditIcon from '@mui/icons-material/Edit';
import SaveIcon from '@mui/icons-material/Save';
import { Box, IconButton, TextField, Typography } from '@mui/material';
import React, { useState } from 'react';

type CharacterBioProps = {
  name: string;
  age: number;
  level: number;
  description: string;
  enableEdit: boolean;
};

const CharacterBio = ({ name, age, level, description, enableEdit }: CharacterBioProps) => {
  const [character, setCharacter] = useState({
    name: name,
    age: age,
    level: level,
    description: description,
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
        {enableEdit && (
          <IconButton onClick={() => toggleEditField('name')}>
            {editField.name ? <SaveIcon /> : <EditIcon />}
          </IconButton>
        )}
      </Box>

      {/* Age Field */}
      <Box display="flex" alignItems="center" mb={2} justifyContent="space-between" border="1px solid #ccc" borderRadius={3} padding={1}>
        {editField.age ? (
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
        {enableEdit && (
          <IconButton onClick={() => toggleEditField('age')}>
            {editField.age ? <SaveIcon /> : <EditIcon />}
          </IconButton>
        )}
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
        {enableEdit && (
          <IconButton onClick={() => toggleEditField('level')}>
            {editField.level ? <SaveIcon /> : <EditIcon />}
          </IconButton>
        )}
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
        {enableEdit && (
          <IconButton onClick={() => toggleEditField('description')}>
            {editField.description ? <SaveIcon /> : <EditIcon />}
          </IconButton>
        )}
      </Box>
    </Box>
  );
};

export default CharacterBio;
