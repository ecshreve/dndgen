import { useQuery } from "@apollo/client";
import { FormControl, InputLabel, MenuItem, Select, Typography } from "@mui/material";
import React, { useState } from "react";
import { RaceDetailsFragment } from "../../__generated__/graphql";
import { GET_RACE_DETAILS } from "../../queries/getRaceDetails";

interface RacePickerProps {
    onSelect: (raceDetails: RaceDetailsFragment) => void;
}

const RacePicker = ({onSelect}: RacePickerProps) => {
    const { data, loading, error } = useQuery(GET_RACE_DETAILS);
    const [selectedRace, setSelectedRace] = useState<string | null>(null);
    const [selectedRaceDetails, setSelectedRaceDetails] = useState<RaceDetailsFragment | null>(null);

    const handleChange = (event: any) => {
        setSelectedRace(event.target.value as string);
        const raceDetails = data.races.find((cls: RaceDetailsFragment) => cls.id === event.target.value);
        setSelectedRaceDetails(raceDetails);
        onSelect(raceDetails);
    };

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;

    return (
        <>
        <FormControl>
            <InputLabel id="race-label">Race</InputLabel>
            <Select
                labelId="race-label"
                id="race-select"
                name="race"
                label="Race"
                value={selectedRace}
                onChange={handleChange}
            >
                {data.races.map((cls: { id: string, name: string }) => (
                    <MenuItem key={cls.id} value={cls.id}>{cls.name}</MenuItem>
                ))}
            </Select>
        </FormControl>
        
    
                      
        <Typography variant="subtitle1" gutterBottom>
          <strong>Proficiencies:</strong>{" "}
          <ul>
          {selectedRaceDetails?.startingProficiencies!.map((prof: any) => <li>{prof.indx}</li>)}
          </ul>
        </Typography>
        </>
    );
};

export default RacePicker;