# APIs

GET /health

POST /api/v1/sign { document, signature, publicKey } -> 
    Sign a document and save for that 

GET /api/v1/signed/:publicKey/:timestamp/:signedTimestam -> 
    returns all signed documents for a given publicKey (timestamp can't be holder that 
    60 seconds and signedTimestamp need to be timestamp signed with the corresponding 
    privateKey

POST /api/v1/subscription { document, name, surname, publicKey, signature } -> 
    Create a new user for the given public Key.
    - publicKey cannot be replicated
    - signature need to be created from document and privateKey for that publicKey
    - name and surname not null



