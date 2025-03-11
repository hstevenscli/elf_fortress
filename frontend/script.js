// Make an instance of two and place it on the page.
var params = {
    fullscreen: true
  };

var elem = document.body;
var two = new Two(params).appendTo(elem);

//Draw rectangle
var x = two.width * 0.5;
var y = two.height * 0.5;
var width = 50;
var height = 50;
var rect = two.makeRectangle(x, y, width, height);  
rect.fill = 'rgb(0, 200, 255)';
rect.opacity = 0.75;
rect.noStroke();

// Cant get the sprite to load for some reason
// let sprite = new Two.Sprite('assets/Elf.png', two.width / 2, two.height / 2, 1, 1, 30);    
// sprite.scale = 0.5; // Adjust size
// two.add(sprite); // Add to the scene

let speed = 3;
let keys = {}; // Object to track pressed keys

// Event listeners for key press and release
window.addEventListener('keydown', (event) => keys[event.key] = true);
window.addEventListener('keyup', (event) => delete keys[event.key]);

// Animation loop
two.bind('update', function() {
    if (keys['ArrowUp']) rect.translation.y -= speed;
    if (keys['ArrowDown']) rect.translation.y += speed;
    if (keys['ArrowLeft']) rect.translation.x -= speed;
    if (keys['ArrowRight']) rect.translation.x += speed;
});

// Start animation loop
two.play();