import { useQuery } from '@apollo/client';
import {
  Box,
  CircularProgress,
  FormControl,
  InputLabel,
  MenuItem,
  Paper,
  Select,
  Stack,
  Typography,
} from '@mui/material';
import React, { useState } from 'react';
import { GET_CLASSES } from './queries/getClasses';

interface ClassDetails {
  id: string;
  indx: string;
  name: string;
  savingThrows: { indx: string }[];
  proficiencies: { indx: string }[];
  proficiencyOptions: { desc: string }[];
  startingEquipment: { quantity: number; equipment: { indx: string } }[];
}

const App: React.FC = () => {
  const [selectedClass, setSelectedClass] = useState<string | null>(null);

  const { loading, error, data } = useQuery(GET_CLASSES);

  const handleClassChange = (event: any) => {
    setSelectedClass(event.target.value as string);
  };

  if (loading) return <CircularProgress />;
  if (error) return <p>Error loading data...</p>;

  const classes = data.classes.edges.map((edge: { node: ClassDetails }) => edge.node);

  const selectedClassDetails = classes.find((cls: ClassDetails) => cls.indx === selectedClass);

  return (
    <Box p={2}>
      <Stack direction="row" spacing={2}>
        {/* Left column: Form */}
        <Box flex={1}>
          <FormControl fullWidth>
            <InputLabel id="class-select-label">Select Class</InputLabel>
            <Select
              labelId="class-select-label"
              value={selectedClass ?? ''}
              onChange={handleClassChange}
              label="Select a class"
            >
              {classes.map((cls: ClassDetails) => (
                <MenuItem key={cls.id} value={cls.indx}>
                  {cls.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Box>

        {/* Right column: Details pane */}
        <Box flex={1}>
          {selectedClassDetails ? (
            <Paper elevation={3}>
              <Box p={2}>
                <Typography variant="h5" gutterBottom>
                  {selectedClassDetails.name}
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                  <strong>Saving Throws:</strong> {selectedClassDetails.savingThrows.map((st: { indx: any; }) => st.indx).join(', ')}
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                  <strong>Proficiencies:</strong> {selectedClassDetails.proficiencies.map((p: { indx: any; }) => p.indx).join(', ')}
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                  <strong>Proficiency Options:</strong> {selectedClassDetails.proficiencyOptions.map((po: { desc: any; }) => po.desc).join(', ')}
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                  <strong>Starting Equipment:</strong>
                  <ul>
                    {selectedClassDetails.startingEquipment.map((eq: { quantity: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; equipment: { indx: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; }; }, idx: React.Key | null | undefined) => (
                      <li key={idx}>
                        {eq.quantity}x {eq.equipment.indx}
                      </li>
                    ))}
                  </ul>
                </Typography>
              </Box>
            </Paper>
          ) : (
            <Typography variant="subtitle1" color="textSecondary">
              Please select a class to view the details.
            </Typography>
          )}
        </Box>
      </Stack>
    </Box>
  );
};

export default App;
