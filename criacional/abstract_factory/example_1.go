package abstractfactory

import "fmt"

// Botao representa um botão na interface gráfica
type Botao interface {
	Render() string
}

// Janela representa uma janela na interface gráfica
type Janela interface {
	Render() string
}

// BotaoWindows é um botão no estilo Windows
type BotaoWindows struct{}

func (b BotaoWindows) Render() string {
	return "Renderizando Botão no estilo Windows"
}

// BotaoMacOS é um botão no estilo MacOS
type BotaoMacOS struct{}

func (b BotaoMacOS) Render() string {
	return "Renderizando Botão no estilo MacOS"
}

// JanelaWindows é uma janela no estilo Windows
type JanelaWindows struct{}

func (j JanelaWindows) Render() string {
	return "Renderizando Janela no estilo Windows"
}

// JanelaMacOS é uma janela no estilo MacOS
type JanelaMacOS struct{}

func (j JanelaMacOS) Render() string {
	return "Renderizando Janela no estilo MacOS"
}

// GUIFactory é a Abstract Factory que cria botões e janelas
type GUIFactory interface {
	CriarBotao() Botao
	CriarJanela() Janela
}

// WindowsFactory cria elementos no estilo Windows
type WindowsFactory struct{}

func (w WindowsFactory) CriarBotao() Botao {
	return BotaoWindows{}
}

func (w WindowsFactory) CriarJanela() Janela {
	return JanelaWindows{}
}

// MacOSFactory cria elementos no estilo MacOS
type MacOSFactory struct{}

func (m MacOSFactory) CriarBotao() Botao {
	return BotaoMacOS{}
}

func (m MacOSFactory) CriarJanela() Janela {
	return JanelaMacOS{}
}

// CriarFactory retorna a fábrica correta com base no sistema operacional
func CriarFactory(sistema string) GUIFactory {
	if sistema == "windows" {
		return WindowsFactory{}
	} else if sistema == "macos" {
		return MacOSFactory{}
	}
	return nil
}

func main() {
	sistema := "macos" // Pode ser "windows" ou "macos"

	fabrica := CriarFactory(sistema)

	botao := fabrica.CriarBotao()
	janela := fabrica.CriarJanela()

	fmt.Println(botao.Render())
	fmt.Println(janela.Render())
}
