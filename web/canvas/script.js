const canvas = document.getElementById("canvas1");
const ctx = canvas.getContext("2d");
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

function drawCircle() {
  ctx.fillStyle = "blue";
  ctx.beginPath();
  ctx.arc(mousePosition.x, mousePosition.y, 3, 0, Math.PI * 2);
  ctx.fill();
}

class Particle {
  constructor(){
    self.x = mousePosition.x;
    self.y = mousePosition.y;
    size = Math.random() * 5 + 1;
  }
}

let animate = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  drawCircle();
  requestAnimationFrame(animate);
}

animate();
