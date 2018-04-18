//import QtQuick.Window 2.2
import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import Qt.labs.settings 1.0
import Client 1.0


ApplicationWindow {
    id: window
    width: 360
    maximumWidth: 360
    height: 520
    visible: true
    title: "Sis. Monitoramento Residencial"

//    property string style: "Material"

    Material.theme: Material.Dark
    Material.primary: Material.Blue
    Material.accent: Material.LightBlue

    Client {
      id: client
      state: 0
      data: ({})
    }

    header: ToolBar {
        id: toolbar
        visible: false
        Material.foreground: "white"

        RowLayout {
            spacing: 20
            anchors.fill: parent

            Label {
                id: titleLabel
                text: "SMCRA"
                font.pixelSize: 20
//                elide: Label.ElideRight
                horizontalAlignment: Qt.AlignHCenter
                verticalAlignment: Qt.AlignVCenter
                Layout.fillWidth: true
            }

            ToolButton {
                contentItem: Image {
                    fillMode: Image.Pad
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Image.AlignVCenter
                    source: "qrc:///qml/images/menu.png"
                }
                onClicked: optionsMenu.open()

                Menu {
                    id: optionsMenu
                    x: parent.width - width
                    transformOrigin: Menu.TopRight


                    MenuItem {
                        id: logoutMenuItem
                        text: "Sair"
                        onTriggered: Controller.logout()
                    }
                }
            }
        }
    }


    StackView {
        id: stackView
        anchors.fill: parent

        initialItem: Pane {
          id: loginPane
          anchors.fill: parent

          ColumnLayout {
            width: 320
            anchors.centerIn: parent
            anchors.verticalCenterOffset: -50
            spacing: 3

            TextField {
               id: login
               Layout.fillWidth: true
               placeholderText: "Digite seu CPF"
               enabled: true
            }

             Button {
                 id: proccessButton

                 Layout.fillWidth: true
                 spacing: 1
                 highlighted: true
                 Material.background: Material.color(Material.Blue)
                 onClicked: Controller.login(login.text)
             }


             TextArea {
                 id: data
                 text: "Não Logado\n\n"
                 readOnly: true
//                 Layout.fillHeight: true
                 width: parent.width

             }

              Connections {
                target: Controller
                onSessionAuthenticated: {
                  client.state = 1

                  var JsonObject = JSON.parse(reply);

                  client.data = JsonObject

                  data.text = "Usuário '"+ login.text +"' Autenticado \n\n"

                  stackView.push("qrc:///qml/pages/SwipePage.qml")

                }
                onSessionAuthenticationError: {
                   client.state = 0
                   data.text = "Autenticação do usuároio '"+ login.text +"' falhou.\n\n"
                }
                onSessionTerminated: {
                  login.text = ""
                  client.state = 0
                  data.text = "Sessão Encerrada.\n\n"
                  stackView.pop()
                }

                onSessionLoader: {
                  client.state = 2
                }
              }

             states: [
                 State {
                     name: "NotAuthenticated"
                     when: client.state == 0
                     PropertyChanges {
                         target: proccessButton
                         text: "Login"
                     }
                     PropertyChanges {
                        target: toolbar
                        visible: false
                     }
                 },

                 State {
                     name: "Authenticated"
                     when: client.state == 1
                     PropertyChanges {
                         target: toolbar
                         visible: true
                     }

                 },

                 State {
                     name: "Authenticating"
                     when: client.state == 2
                     PropertyChanges {
                         target: proccessButton
                         text: "Autenticando..."
                         enabled: false
                     }
                 }
             ]
          }
        }
    }


}
