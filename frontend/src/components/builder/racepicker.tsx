import { useQuery } from "@apollo/client";
import {
    FormControl,
    InputLabel,
    MenuItem,
    Select,
    Typography,
} from "@mui/material";
import React, { useState } from "react";
import { RaceDetailsFragment } from "../../__generated__/graphql";
import { GET_RACE_DETAILS } from "../../queries/getRaceDetails";
import { OptionSelect } from "./optionselect";

interface RacePickerProps {
  onSelect: (raceDetails: RaceDetailsFragment) => void;
  onProficiencyOptionsChange: (options: any[]) => void;
}

const RacePicker = ({ onSelect, onProficiencyOptionsChange }: RacePickerProps) => {
  const { data, loading, error } = useQuery(GET_RACE_DETAILS);
  const [selectedRace, setSelectedRace] = useState<string>("");
  const [selectedRaceDetails, setSelectedRaceDetails] =
    useState<RaceDetailsFragment | null>(null);

  const handleChange = (event: any) => {
    setSelectedRace(event.target.value as string);
    const raceDetails = data.races.find(
      (rr: RaceDetailsFragment) => rr.id === event.target.value
    );
    setSelectedRaceDetails(raceDetails);
    onSelect(raceDetails);
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  return (
    <>
      <FormControl>
        <InputLabel id="race-select-label">Race</InputLabel>
        <Select
          labelId="race-select-label"
          id="race-select"
          label="Race"
          value={selectedRace}
          onChange={handleChange}
        >
          {data.races.map((cls: { id: string; name: string }) => (
            <MenuItem key={cls.id} value={cls.id}>
              {cls.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      {selectedRaceDetails && (
        <>
          <Typography variant="subtitle1" gutterBottom>
            <strong>Proficiencies:</strong>{" "}
            <ul>
              {selectedRaceDetails.startingProficiencies?.map((prof: any) => (
                <li key={prof.id}>{prof.indx}</li>
              ))}
            </ul>
          </Typography>

          {selectedRaceDetails?.startingProficiencyOptions!.length > 0 && (
            <Typography variant="subtitle1" gutterBottom>
              <strong>Proficiency Options:</strong>{" "}
              {selectedRaceDetails.startingProficiencyOptions!.map(
                (spo: any) => (
                  <OptionSelect
                    key={spo.id}
                    label={spo.desc || spo.choose}
                    choose={spo.choose}
                    options={spo.proficiencies || []}
                    handleSelect={onProficiencyOptionsChange}
                  />
                )
              )}
            </Typography>
          )}
        </>
      )}
    </>
  );
};

export default RacePicker;
