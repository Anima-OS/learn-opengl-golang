package getstarted

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/raedatoui/learn-opengl-golang/sections"
	"github.com/raedatoui/learn-opengl-golang/utils"
)

type HelloCoordinates struct {
	sections.BaseSketch
	shader             uint32
	vao, vbo, ebo      uint32
	texture1, texture2 uint32
	transform          mgl32.Mat4
	cubePositions      []mgl32.Vec3
}

func (hc *HelloCoordinates) Setup(w *glfw.Window, f *utils.Font) error {
	hc.Name = "6. Coordinate Systems"
	hc.Window = w
	hc.Font = f
	hc.Color = utils.StepColor(utils.MAG, utils.BLACK, 10, 7)

	var err error
	hc.shader, err = utils.Shader("_assets/6.coordinates/coordinate.vs",
		"_assets/6.coordinates/coordinate.frag", "")
	if err != nil {
		return err
	}
	gl.UseProgram(hc.shader)

	vertices := []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	hc.cubePositions = []mgl32.Vec3{
		mgl32.Vec3{0.0, 0.0, 0.0},
		mgl32.Vec3{2.0, 5.0, -15.0},
		mgl32.Vec3{-1.5, -2.2, -2.5},
		mgl32.Vec3{-3.8, -2.0, -12.3},
		mgl32.Vec3{2.4, -0.4, -3.5},
		mgl32.Vec3{-1.7, 3.0, -7.5},
		mgl32.Vec3{1.3, -2.0, -2.5},
		mgl32.Vec3{1.5, 2.0, -2.5},
		mgl32.Vec3{1.5, 0.2, -1.5},
		mgl32.Vec3{-1.3, 1.0, -1.5},
	}

	gl.GenVertexArrays(1, &hc.vao)
	gl.GenBuffers(1, &hc.vbo)
	gl.GenBuffers(1, &hc.ebo)

	gl.BindVertexArray(hc.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, hc.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*utils.GL_FLOAT32_SIZE, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*utils.GL_FLOAT32_SIZE, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// TexCoord attribute
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 5*utils.GL_FLOAT32_SIZE, gl.PtrOffset(3*utils.GL_FLOAT32_SIZE))
	gl.EnableVertexAttribArray(2)

	gl.BindVertexArray(0) // Unbind VAO

	// ====================
	// Texture 1
	// ====================
	gl.GenTextures(1, &hc.texture1)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture1)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	rgba, err := utils.ImageToPixelData("_assets/images/container.png")
	if err != nil {
		panic(err)
	}
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	// ====================
	// Texture 2
	// ====================
	gl.GenTextures(1, &hc.texture2)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture2)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	rgba, err = utils.ImageToPixelData("_assets/images/awesomeface.png")
	if err != nil {
		return err
	}
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return nil
}

func (hc *HelloCoordinates) Update() {

}

func (hc *HelloCoordinates) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(hc.Color.R, hc.Color.G, hc.Color.B, hc.Color.A)

	// Bind Textures using texture units
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture1)
	loc1 := gl.GetUniformLocation(hc.shader, gl.Str("ourTexture1\x00"))
	gl.Uniform1i(loc1, 0)

	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture2)
	loc2 := gl.GetUniformLocation(hc.shader, gl.Str("ourTexture2\x00"))
	gl.Uniform1i(loc2, 1)

	// Activate shader
	gl.UseProgram(hc.shader)

	// Create transformations
	view := mgl32.Translate3D(0.0, 0.0, -3.0)
	projection := mgl32.Perspective(45.0, 800.0/600.0, 0.1, 100.0)
	// Get their uniform location
	modelLoc := gl.GetUniformLocation(hc.shader, gl.Str("model\x00"))
	viewLoc := gl.GetUniformLocation(hc.shader, gl.Str("view\x00"))
	projLoc := gl.GetUniformLocation(hc.shader, gl.Str("projection\x00"))
	// Pass the matrices to the shader
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
	// Note: currently we set the projection matrix each frame,
	// but since the projection matrix rarely changes it's often best practice to set it outside the main loop only once.
	gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])

	// Draw container
	gl.BindVertexArray(hc.vao)

	for i := 0; i < 10; i++ {
		// Calculate the model matrix for each object and pass it to shader before drawing
		model := mgl32.Translate3D(
			hc.cubePositions[i][0],
			hc.cubePositions[i][1],
			hc.cubePositions[i][2])
		//angle := 20.0 * float32(i)
		//if i % 3 == 0 {
		angle := float32(glfw.GetTime()) * float32(i+1)
		//}  // Every 3rd iteration (including the first) we set the angle using GLFW's time function.

		model = model.Mul4(mgl32.HomogRotate3D(angle, mgl32.Vec3{1.0, 0.3, 0.5}))
		gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}
	gl.BindVertexArray(0)

	hc.Font.SetColor(0.0, 0.0, 0.0, 1.0)
	hc.Font.Printf(30, 30, 0.5, hc.Name)
}

func (hc *HelloCoordinates) Close() {
	gl.DeleteVertexArrays(1, &hc.vao)
	gl.DeleteBuffers(1, &hc.vbo)
	gl.DeleteBuffers(1, &hc.ebo)
	gl.UseProgram(0)
}

func (hc *HelloCoordinates) HandleKeyboard(k glfw.Key, s int, a glfw.Action, mk glfw.ModifierKey) {

}

func (hc *HelloCoordinates) HandleMousePosition(xpos, ypos float64) {

}

func (hc *HelloCoordinates) HandleScroll(xoff, yoff float64) {

}
