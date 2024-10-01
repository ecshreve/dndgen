import { useQuery } from "@apollo/client";
import { ExpandMore } from "@mui/icons-material";
import { Accordion, AccordionDetails, AccordionSummary, Box, FormControl, InputLabel, MenuItem, Select, Typography } from "@mui/material";
import React, { useState } from "react";
import { ClassDetailsFragment } from "../../__generated__/graphql";
import { GET_CLASS_DETAILS } from "../../queries/getClassDetails";

const ClassPicker = () => {
    const { data, loading, error } = useQuery(GET_CLASS_DETAILS);
    const [selectedClass, setSelectedClass] = useState<string | null>(null);
    const [selectedClassDetails, setSelectedClassDetails] = useState<ClassDetailsFragment | null>(null);
   
    const handleChange = (event: any) => {
        setSelectedClass(event.target.value as string);
        const classDetails = data.classes.find((cls: ClassDetailsFragment) => cls.id === event.target.value);
        setSelectedClassDetails(classDetails);
    };

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    return (
        <>
        <FormControl fullWidth>
            <InputLabel id="class-label">Class</InputLabel>
            <Select
                labelId="class-label"
                id="class-select"
                name="class"
                label="Class"
                value={selectedClass}
                onChange={handleChange}
            >
                {data.classes.map((cls: { id: string, name: string }) => (
                    <MenuItem key={cls.id} value={cls.id}>{cls.name}</MenuItem>
                ))}
            </Select>
        </FormControl>
        {selectedClassDetails && (
              <Box flex={1} margin="normal">
              <Accordion>
                <AccordionSummary
                  expandIcon={selectedClassDetails && <ExpandMore />}
                  aria-controls="class-content"
                  id="class-header"
                >
                  <Typography variant="h6">
                    {selectedClassDetails
                      ? selectedClassDetails.name + " Details"
                      : "Select a class"}
                  </Typography>
                </AccordionSummary>
                  <AccordionDetails>
                    <Box p={2}>
                      <Typography variant="h5" gutterBottom>
                        {selectedClassDetails.name}
                      </Typography>
    
                      
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Proficiencies:</strong>{" "}
                      {selectedClassDetails.proficiencies!.map((prof: any) => prof.indx)}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Proficiency Options:</strong>{" "}
                      {selectedClassDetails.proficiencyOptions!.map((po: any) => po.desc)}
                    </Typography>
                    
                    </Box>
                  </AccordionDetails>
                </Accordion>
              </Box>
        )}
        </>
    );
};

export default ClassPicker;