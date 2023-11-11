# Utenti

* Ogni utente è in grado di effettuare richiesta per multiple credenziali.
* Sull'interfaccia non è possibile richiedere un'altra credenziale se ce ne è già una approvata (anche se tecnicamente possibile)

# Signable

* Un documento da firmare è assegnato ad una coppia (codice fiscale, email)
* Il documento firmato sarà poi assegnato alla credenziale che è stata utilizzata per firmare il documento.
* Un documento può essere firmato se:
    - il codice fiscale e email coincidono con quelle della credenziale usata
    - la signature del documento coincide con la chiave pubblica utilizzata per firmare 

