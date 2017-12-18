# ocmr

> Open Computable Medical Record


The OCMR is a simple snapshot of clinical information designed to provide important contextual information in relation to human and computer decision making.

It is designed as a data structure:

* to create public repositories of machine learning training data
* as an intermediary data structure created from electronic health records to encapsulate a clinical scenario to be passed to an algorithm for processing.


## Example v1

```json
  {
    "Age": 73,
    "Sex": "male",
    "Problems": [
      {
        "Concept": {
          "ConceptID": 29857009,
          "FullySpecifiedName": "Chest pain (finding)"
        },
        "Duration": "Acute"
      },
      {
        "Concept": {
          "ConceptID": 267036007,
          "FullySpecifiedName": "Dyspnea (finding)"
        },
        "Duration": "Acute"
      },
      {
        "Concept": {
          "ConceptID": 415690000,
          "FullySpecifiedName": "Sweating (finding)"
        },
        "Duration": "Acute"
      },
      {
        "Concept": {
          "ConceptID": 426555006,
          "FullySpecifiedName": "Pain radiating to jaw (finding)"
        },
        "Duration": "Acute"
      },
      {
        "Concept": {
          "ConceptID": 76388001,
          "FullySpecifiedName": "ST segment elevation (finding)"
        },
        "Duration": "Acute"
      }
    ],
    "Answer": {
      "ConceptID": 22298006,
      "FullySpecifiedName": "Myocardial infarction (disorder)"
    }
  }
```

Here, a middle-aged male patient presents with chest pain, breathlessness (dyspnoea), sweating, pain radiating to the jaw and ECG changes showing ST elevation; the diagnosis is Myocardial infarction (heart attack). 

For version 1, each problem simply has an associated SNOMED-CT concept and duration. It is conceivable that additional information could be recorded, even imaging data, with the OCMR acting as a contextual wrapper around that data.


