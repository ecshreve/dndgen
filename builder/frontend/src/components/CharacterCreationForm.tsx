// CharacterCreationForm.tsx

import React, { useState } from 'react';
import './CharacterCreationForm.css';
import { CharacterFormValues, Option } from './types';

interface CharacterCreationFormProps {
  classes: Option[];
  races: Option[];
  alignments: Option[];
  onSubmit: (values: CharacterFormValues) => void;
}

const CharacterCreationForm: React.FC<CharacterCreationFormProps> = ({
  classes,
  races,
  alignments,
  onSubmit,
}) => {
  const [formValues, setFormValues] = useState<CharacterFormValues>({
    name: '',
    classId: '',
    raceId: '',
    alignment: '',
    attributes: {
      strength: 8,
      dexterity: 8,
      constitution: 8,
      intelligence: 8,
      wisdom: 8,
      charisma: 8,
    },
  });

  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;

    setFormValues((prevValues) => ({
      ...prevValues,
      [name]: value,
    }));
  };

  const handleAttributeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setFormValues((prevValues) => ({
      ...prevValues,
      attributes: {
        ...prevValues.attributes,
        [name]: Number(value),
      },
    }));
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const validationErrors: { [key: string]: string } = {};

    if (!formValues.name.trim()) {
      validationErrors.name = 'Name is required';
    }
    if (!formValues.classId) {
      validationErrors.classId = 'Class is required';
    }
    if (!formValues.raceId) {
      validationErrors.raceId = 'Race is required';
    }
    if (!formValues.alignment) {
      validationErrors.alignment = 'Alignment is required';
    }

    setErrors(validationErrors);

    if (Object.keys(validationErrors).length === 0) {
      onSubmit(formValues);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="form-section">
        <div className="form-group">
          <label htmlFor="name">Character Name</label>
          <input
            id="name"
            name="name"
            type="text"
            value={formValues.name}
            onChange={handleChange}
            />
          {errors.name && <div className="error">{errors.name}</div>}
        </div>

        <div className="form-group">
        <label htmlFor="classId">Class</label>
        <select
          id="classId"
          name="classId"
          value={formValues.classId}
          onChange={handleChange}
        >
          <option value="">Select Class</option>
          {classes.map((cls) => (
            <option key={cls.id} value={cls.id}>
              {cls.name}
            </option>
          ))}
        </select>
        {errors.classId && <div className="error">{errors.classId}</div>}
      </div>

      <div className="form-group">
        <label htmlFor="raceId">Race</label>
        <select
          id="raceId"
          name="raceId"
          value={formValues.raceId}
          onChange={handleChange}
        >
          <option value="">Select Race</option>
          {races.map((race) => (
            <option key={race.id} value={race.id}>
              {race.name}
            </option>
          ))}
        </select>
        {errors.raceId && <div className="error">{errors.raceId}</div>}
      </div>

      <div className="form-group">
        <label htmlFor="alignment">Alignment</label>
        <select
          id="alignment"
          name="alignment"
          value={formValues.alignment}
          onChange={handleChange}
        >
          <option value="">Select Alignment</option>
          {alignments.map((alignment) => (
            <option key={alignment.id} value={alignment.id}>
              {alignment.name}
            </option>
          ))}
        </select>
        {errors.alignment && <div className="error">{errors.alignment}</div>}
      </div>
    </div>

    <div>
      <h3>Ability Scores</h3>
    </div>
      <div className="form-section">
        {(['strength', 'dexterity', 'constitution', 'intelligence', 'wisdom', 'charisma'] as const).map(
          (attr) => (
            <div className="form-group" key={attr}>
              <label htmlFor={attr}>{attr.charAt(0).toUpperCase() + attr.slice(1)}</label>
              <input
                id={attr}
                name={attr}
                type="number"
                min="1"
                max="20"
                value={formValues.attributes[attr]}
                onChange={handleAttributeChange}
              />
            </div>
          )
        )}
      </div>

      <button type="submit">Create Character</button>
    </form>
  );
};

export default CharacterCreationForm;
