import { FormControl, InputLabel, MenuItem, Select, SelectChangeEvent, Typography } from "@mui/material";
import React, { useState } from "react";

interface OptionSelectProps {
    label: string;
    choose: number;
    options: { id: string; indx: string, selected: boolean }[];
    handleSelect: (options: string[]) => void;
}

export const OptionSelect = ({ label, choose, options, handleSelect }: OptionSelectProps) => {
    const [selectedOption, setSelectedOption] = useState<string[]>([]);
    
    const handleChange = (event: SelectChangeEvent<typeof selectedOption>) => {
        const {
          target: { value },
        } = event;
        setSelectedOption(
          // On autofill we get a stringified value.
          typeof value === 'string' ? value.split(',') : value,
        );
        handleSelect(typeof value === 'string' ? value.split(',') : value);
      };

    return (
        <>
        {options.length > 0 ? (
        <FormControl fullWidth>
            <InputLabel id={`${label}-select-label`}>{label}</InputLabel>
            <Select
                labelId={`${label}-select-label`}
                value={selectedOption}
                label={label}
                multiple={choose > 1}
                onChange={handleChange}
            >
                {options.map((option: { id: string; indx: string }) => (
                    <MenuItem key={option.id} value={option.indx}>
                        {option.indx}
                    </MenuItem>
                ))}
                </Select>
            </FormControl>
        ) : (
            <Typography variant="body1">No options available</Typography>
        )}
        </>
    );
};
