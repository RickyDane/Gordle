let answerWord;
let currentRow = 0;
let currentTile = 0;
let guess = '';

const board = document.getElementById('game-board');
const keys = document.querySelectorAll('.key');
let allWords = [];

// Spielfeld initialisieren
async function initBoard() {
    // Alle Wörter laden
    allWords = (await fetch("/getAllGermanWords").then(response => response.text())).split("\n").map(word => word.toLowerCase());
    console.time("allWordsLoaded");
    console.timeEnd("allWordsLoaded");
    console.log(allWords);
    // 6 Reihen, 5 Spalten
    for (let i = 0; i < 40; i++) {
        const tile = document.createElement('div');
        tile.classList.add('tile');
        tile.setAttribute('id', `tile-${i}`);
        board.appendChild(tile);
    }
    console.time("boardCreated");
    while (answerWord?.length != 5 || answerWord?.includes("ä") || answerWord?.includes("ö") || answerWord?.includes("ü")) {
        answerWord = allWords[Math.floor(Math.random() * allWords.length)];
    }
    console.timeEnd("boardCreated");
}

// Eingabe von Tasten
keys.forEach(key => {
    key.addEventListener('click', () => {
        const keyLetter = key.innerText;
        if (keyLetter === 'Enter') {
            submitGuess();
        } else if (keyLetter === 'Löschen') {
            deleteLetter();
        } else {
            addLetter(keyLetter);
        }
    });
});

// Eingabe von Tasten mit Tastendruck
document.addEventListener('keydown', (event) => {
    const key = event.key;
    if (key === 'Enter') {
        submitGuess();
    } else if (key === 'Backspace' || key === 'Delete') {
        deleteLetter();
    } else if (key.length === 1 && key.match(/[a-z]/i)) {
        addLetter(key);
    }
});

// Buchstaben hinzufügen
function addLetter(letter) {
    if (currentTile < 5) {
        const tile = document.getElementById(`tile-${currentRow * 5 + currentTile}`);
        tile.innerText = letter.toUpperCase();
        guess += letter;
        currentTile++;
    }
}

// Buchstaben löschen
function deleteLetter() {
    if (currentTile > 0) {
        currentTile--;
        const tile = document.getElementById(`tile-${currentRow * 5 + currentTile}`);
        tile.innerText = '';
        guess = guess.slice(0, -1);
    }
}

// Rateversuch überprüfen
function submitGuess() {
    if (guess.length === 5) {
        if (!allWords.includes(guess.toLowerCase())) {
            alert("Das Wort ist ungültig.");
            return;
        }
        checkGuess();
        currentRow++;
        currentTile = 0;
        guess = '';
    } else {
        alert("Das Wort muss 5 Buchstaben haben.");
    }
}

// Rateversuch validieren
function checkGuess() {
    const guessArray = guess.split('');
    for (let i = 0; i < 5; i++) {
        const tile = document.getElementById(`tile-${currentRow * 5 + i}`);
        const letter = guessArray[i];

        if (answerWord[i] === letter) {
            tile.classList.add('correct');
            keys.forEach(key => {
               if (key.innerText.toLowerCase() === letter) {
                   key.classList.add('correct');
               } 
            });
        } else if (answerWord.includes(letter)) {
            tile.classList.add('present');
            keys.forEach(key => {
               if (key.innerText.toLowerCase() === letter) {
                   key.classList.add('present');
               } 
            });
        } else {
            tile.classList.add('absent');
            keys.forEach(key => {
               if (key.innerText.toLowerCase() === letter) {
                   key.classList.add('absent');
               } 
            });
        }
    }

    if (guess === answerWord) {
        alert("Glückwunsch! Du hast das Wort erraten.");
        location.href = 'https://google.com/search?q=' + answerWord;
    } else if (currentRow === 7) { // 8. Rateversuch
        alert(`Das richtige Wort war: ${answerWord}`);
        location.href = 'https://google.com/search?q=' + answerWord;
    }
    console.log(currentRow);
}

// Initialisiere das Spielfeld
initBoard();