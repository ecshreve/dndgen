import { useQuery } from "@apollo/client";
import { FormControl, InputLabel, MenuItem, Select, Typography } from "@mui/material";
import React, { useState } from "react";
import { ClassDetailsFragment } from "../../__generated__/graphql";
import { GET_CLASS_DETAILS } from "../../queries/getClassDetails";

interface ClassPickerProps {
    onSelect: (classDetails: ClassDetailsFragment) => void;
}

const ClassPicker = ({onSelect}: ClassPickerProps) => {
    const { data, loading, error } = useQuery(GET_CLASS_DETAILS);
    const [selectedClass, setSelectedClass] = useState<string | null>(null);
    const [selectedClassDetails, setSelectedClassDetails] = useState<ClassDetailsFragment | null>(null);

    const handleChange = (event: any) => {
        setSelectedClass(event.target.value as string);
        const classDetails = data.classes.find((cls: ClassDetailsFragment) => cls.id === event.target.value);
        setSelectedClassDetails(classDetails);
        onSelect(classDetails);
    };

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    return (
        <>
        <FormControl>
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
        
    
                      
        <Typography variant="subtitle1" gutterBottom>
          <strong>Proficiencies:</strong>{" "}
          {selectedClassDetails?.proficiencies!.map((prof: any) => prof.indx)}
        </Typography>
        <Typography variant="subtitle1" gutterBottom>
          <strong>Proficiency Options:</strong>{" "}
          {selectedClassDetails?.proficiencyOptions!.map((po: any) => po.desc)}
        </Typography>
        </>
    );
};

export default ClassPicker;