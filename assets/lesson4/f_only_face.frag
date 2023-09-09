#version 330 core
out vec4 FragColor;

uniform sampler2D texture1;
uniform sampler2D texture2;

in vec2 TexCoord;

void main()
{
    vec2 rotor = vec2(-1.0, 1.0);
    FragColor = mix(texture(texture1, TexCoord), texture(texture2, TexCoord * rotor), 0.2);
}