# go-gin-poc

Lista dei comandi eseguiti per il deployment su HEROKU:

- heroku local > per eseguire l'app in locale
- heroku create go-video-app > per creare l'app in locale (la prima volta chiede le credenziali)
- git push heroku heroku:master > si deve passare il nome del branch su cui si è e quello del branch su heroku che è sempre MASTER
- heroku apps > elenca le app distribuite
- heroku open > apre nel browser l'app caricata
- heroku logs > mostra i log
- heroku destroy --confirm [nome app] > distrugge la app
