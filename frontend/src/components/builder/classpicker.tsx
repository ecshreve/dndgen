import { useQuery } from "@apollo/client";
import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Typography,
} from "@mui/material";
import React, { useState } from "react";
import { ClassDetailsFragment } from "../../__generated__/graphql";
import { GET_CLASS_DETAILS } from "../../queries/getClassDetails";
import { OptionSelect } from "./optionselect";

interface ClassPickerProps {
  onSelect: (classDetails: ClassDetailsFragment) => void;
  onProficiencyOptionsChange: (options: any[]) => void;
}

const ClassPicker = ({ onSelect, onProficiencyOptionsChange }: ClassPickerProps) => {
  const { data, loading, error } = useQuery(GET_CLASS_DETAILS);
  const [selectedClass, setSelectedClass] = useState<string>("");
  const [selectedClassDetails, setSelectedClassDetails] =
    useState<ClassDetailsFragment | null>(null);

  const handleChange = (event: any) => {
    setSelectedClass(event.target.value as string);
    const classDetails = data.classes.find(
      (cls: ClassDetailsFragment) => cls.id === event.target.value
    );
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
          {data.classes.map((cls: { id: string; name: string }) => (
            <MenuItem key={cls.id} value={cls.id}>
              {cls.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      {selectedClassDetails && (
        <>
          <Typography variant="subtitle1" gutterBottom>
            <strong>Proficiencies:</strong>{" "}
            <ul>
              {selectedClassDetails?.proficiencies!.map((prof: any) => (
                <li key={prof.id}>{prof.indx}</li>
              ))}
            </ul>
          </Typography>

          {selectedClassDetails?.proficiencyOptions && selectedClassDetails?.proficiencyOptions.length > 0 && (
            <Typography variant="subtitle1" gutterBottom>
              <strong>Proficiency Options:</strong>{" "}
              {selectedClassDetails?.proficiencyOptions!.map((po: any) => (
                <OptionSelect
                  key={po.id}
                  label={po.desc || po.choose}
                  choose={po.choose}
                  options={po.proficiencies}
                  handleSelect={onProficiencyOptionsChange}
                />
              ))}
            </Typography>
          )}
        </>
      )}
    </>
  );
};

export default ClassPicker;
