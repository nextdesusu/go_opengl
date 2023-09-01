#version 330 core
layout(location = 0) in vec3 aPos;

out vec4 changedColor;

void main() {
  gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
  changedColor = gl_Position;
}