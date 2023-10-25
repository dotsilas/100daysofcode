const canvas = document.getElementById("canvas1");
const ctx = canvas.getContext("2d");
let ParticlesArray = [];
console.log(ctx);

window.addEventListener("resize", () => {
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
})

const mousePosition = {
  x: undefined,
  y: undefined,
}

canvas.addEventListener("click", (event) => {
  mousePosition.x = event.x;
  mousePosition.y = event.y;
})

canvas.addEventListener("mousemove", (event) => {
  mousePosition.x = event.x;
  mousePosition.y = event.y;
})

class Particle {
  constructor() {
    // this.x = mousePosition.x;
    // this.y = mousePosition.y;
    this.x = Math.random() * canvas.width;
    this.y = Math.random() * canvas.height;
    this.size = Math.random() * 5 + 1;
    this.speedX = Math.random() * 3 - 1.5;
    this.speedY = Math.random() * 3 - 1.5;
  }
  update() {
    this.x = this.speedX;
    this.y = this.speedY;
  }
  draw() {
    ctx.fillStyle = "blue";
    ctx.beginPath();
    ctx.arc(this.x, this.y, 5, 0, Math.PI * 2);
    ctx.fill();
  }
}

function init() {
  for (let i = 0; i < 100; i++) {
    ParticlesArray.push(new Particle());
  }
}
init();

function handleParticles() {
  for (let i = 0; i < ParticlesArray.length; i++) {
    ParticlesArray[i].update();
    ParticlesArray[i].draw();
  }
}

let animate = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  handleParticles();
  requestAnimationFrame(animate);
}

animate();
