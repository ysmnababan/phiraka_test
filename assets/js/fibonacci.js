document.getElementById('fibonacciForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    // Get user input
    const rows = parseInt(document.getElementById('rows').value);
    const columns = parseInt(document.getElementById('columns').value);
    
    // Calculate the number of Fibonacci numbers needed
    const totalNumbers = rows * columns;
    
    // Generate Fibonacci sequence
    const fibonacci = generateFibonacci(totalNumbers);
    
    // Display the Fibonacci grid
    displayFibonacciGrid(fibonacci, rows, columns);
});

function generateFibonacci(n) {
    const fibonacci = [0, 1];
    while (fibonacci.length < n) {
        const nextNumber = fibonacci[fibonacci.length - 1] + fibonacci[fibonacci.length - 2];
        fibonacci.push(nextNumber);
    }
    return fibonacci;
}

function displayFibonacciGrid(fibonacci, rows, columns) {
    const gridContainer = document.getElementById('fibonacciGrid');
    
    // Clear previous grid
    gridContainer.innerHTML = '';
    
    // Set grid template
    gridContainer.style.gridTemplateColumns = `repeat(${columns}, 1fr)`;
    
    // Create and append grid items
    fibonacci.forEach(num => {
        const gridItem = document.createElement('div');
        gridItem.className = 'grid-item';
        gridItem.textContent = num;
        gridContainer.appendChild(gridItem);
    });
}
