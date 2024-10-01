import { useQuery } from "@apollo/client";
import { ExpandMore } from "@mui/icons-material";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  CircularProgress,
  FormControl,
  InputLabel,
  MenuItem,
  Select, Stack, Typography
} from "@mui/material";
import React, { useState } from "react";
import CharacterSummary from "./CharacterSummary";
import { GET_CLASS_DETAILS } from "./queries/getClassDetails";
import { GET_RACE_DETAILS } from "./queries/getRaceDetails";
interface ClassDetails {
    id: string;
    indx: string;
    name: string;
    savingThrows: { 
      id: string;
      indx: string
    }[];
    proficiencies: { 
      id: string;
      indx: string;
      reference: string;
    }[];
    proficiencyOptions: { 
      id: string;
      desc: string;
      choose: number;
      proficiencies: { 
        id: string;
        indx: string;
        name: string;
      }[];
    }[];
    startingEquipment: { 
      id: string;
      quantity: number;
      equipment: {
        id: string;
        indx: string;
      }
    }[];
}

interface RaceDetails {
    id: string;
    indx: string;
    name: string;
    ageDesc: string;
    alignmentDesc: string;
    size: string;
    sizeDesc: string;
    speed: number;
    languageDesc: string;
    startingProficiencies: { 
      id: string;
      indx: string;
      name: string;
    }[];
    startingProficiencyOptions: { 
      id: string;
      indx: string;
      name: string;
      desc: string;
      choose: number;
      proficiencies: { 
        id: string;
        indx: string;
        name: string;
      }[];
    };
  }



const CharacterDropdowns: React.FC = () => {
    const [selectedClass, setSelectedClass] = useState<string | null>(null);
    const [selectedRace, setSelectedRace] = useState<string | null>(null);
  
    // Fetching classes
    const {
      loading: loadingClasses,
      error: errorClasses,
      data: classData,
    } = useQuery(GET_CLASS_DETAILS);
  
    // Fetching races
    const {
      loading: loadingRaces,
      error: errorRaces,
      data: raceData,
    } = useQuery(GET_RACE_DETAILS);
  
    const handleClassChange = (event: any) => {
      setSelectedClass(event.target.value as string);
    };
  
    const handleRaceChange = (event: any) => {
      setSelectedRace(event.target.value as string);
    };
  
    if (loadingClasses || loadingRaces) return <CircularProgress />;
    if (errorClasses || errorRaces)
      return (
        <>
          <p>Error loading data...</p>
          <p>{errorClasses?.message}</p>
          <p>{errorRaces?.message}</p>
        </>
      );
  
    const classes = classData.classes
    const races = raceData.races;
  
    const selectedClassDetails = classes.find(
      (cls: ClassDetails) => cls.indx === selectedClass
    );
    const selectedRaceDetails: RaceDetails = races.find(
      (race: RaceDetails) => race.indx === selectedRace
    );
    console.log(selectedRaceDetails);
  
    return (
      <>
        <CharacterSummary />
        <Box p={2}>
        <Stack direction="row" spacing={2}>
          {/* Left column: Form */}
          <Box flex={1}>
            {/* Race Selection */}
            <FormControl fullWidth margin="normal">
              <InputLabel id="race-select-label">Select Race</InputLabel>
              <Select
                labelId="race-select-label"
                value={selectedRace ?? ""}
                onChange={handleRaceChange}
                label="Select a race"
              >
                {races.map((race: RaceDetails) => (
                  <MenuItem key={race.id} value={race.indx}>
                    {race.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
  
            <FormControl fullWidth>
              <InputLabel id="class-select-label">Select Class</InputLabel>
              <Select
                labelId="class-select-label"
                value={selectedClass ?? ""}
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
          <Box flex={1} margin="normal">
            <Accordion>
              <AccordionSummary
                expandIcon={selectedRaceDetails && <ExpandMore />}
                aria-controls="race-content"
                id="race-header"
              >
                <Typography variant="h6">
                  {selectedRaceDetails
                    ? selectedRaceDetails.name + " Details"
                    : "Select a race"}
                </Typography>
              </AccordionSummary>
              {selectedRaceDetails && (
                <AccordionDetails>
                  <Box p={2}>
                    <Typography variant="h5" gutterBottom>
                      {selectedRaceDetails.name}
                    </Typography>
  
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Age:</strong> {selectedRaceDetails.ageDesc}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Alignment:</strong>{" "}
                      {selectedRaceDetails.alignmentDesc}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Size:</strong> {selectedRaceDetails.size}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Size Description:</strong>{" "}
                      {selectedRaceDetails.sizeDesc}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Speed:</strong> {selectedRaceDetails.speed} feet
                    </Typography>
                    {/* <Typography variant="subtitle1" gutterBottom>
                      <strong>Ability Bonuses:</strong>
                      <ul>
                        {selectedRaceDetails.abilityBonuses.map((ab: { abilityScore: { indx: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; }; bonus: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; }, idx: React.Key | null | undefined) => (
                          <li key={idx}>
                            {ab.abilityScore.indx}: {ab.bonus}
                          </li>
                        ))}
                      </ul>
                    </Typography> */}
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Languages:</strong>{" "}
                      {selectedRaceDetails.languageDesc}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Starting Proficiencies:</strong>
                      <ul>
                        {selectedRaceDetails.startingProficiencies.map(
                          (sp: { indx: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; }, idx: React.Key | null | undefined) => (
                            <li key={idx}>{sp.indx}</li>
                          )
                        )}
                      </ul>
                    </Typography>
                    {selectedRaceDetails.startingProficiencyOptions && (
                      <Typography variant="subtitle1" gutterBottom>
                        <strong>Starting Proficiency Options:</strong>
                        <br />
                        <strong>Description:</strong>{" "}
                        {selectedRaceDetails.startingProficiencyOptions.desc}
                        <br />
                        <strong>Choose:</strong>{" "}
                        {selectedRaceDetails.startingProficiencyOptions.choose}
                        <br />
                        <strong>Proficiencies:</strong>
                        <ul>
                          {selectedRaceDetails.startingProficiencyOptions.proficiencies.map(
                            (spo: { indx: string | number | boolean | React.ReactElement<any, string | React.JSXElementConstructor<any>> | Iterable<React.ReactNode> | React.ReactPortal | null | undefined; }, idx: React.Key | null | undefined) => (
                              <li key={idx}>{spo.indx}</li>
                            )
                          )}
                        </ul>
                      </Typography>
                    )}
                  </Box>
                </AccordionDetails>
              )}
            </Accordion>
  
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
              {selectedClassDetails && (
                <AccordionDetails>
                  <Box p={2}>
                    <Typography variant="h5" gutterBottom>
                      {selectedClassDetails.name}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Saving Throws:</strong>{" "}
                      {selectedClassDetails.savingThrows
                        .map((st: { indx: any }) => st.indx)
                        .join(", ")}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Proficiencies:</strong>{" "}
                      {selectedClassDetails.proficiencies
                        .map((p: { indx: any }) => p.indx)
                        .join(", ")}
                    </Typography>
                    <Typography variant="subtitle1" gutterBottom>
                      <strong>Proficiency Options:</strong>{" "}
                      {selectedClassDetails.proficiencyOptions
                        .map((po: { desc: any }) => po.desc)
                        .join(", ")}
                    </Typography>
                    
                  </Box>
                </AccordionDetails>
              )}
            </Accordion>
          </Box>
        </Stack>
      </Box>
      </>
    );
  };

export default CharacterDropdowns;
