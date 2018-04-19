import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Layouts 1.2
import QtQuick.Controls 1.4
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0


  Pane {
      id: root
      anchors.fill: parent

      ColumnLayout {
          anchors.right: parent.right
          anchors.left: parent.left

          RowLayout {
              Layout.fillWidth: true

              TextField {
                  id: textDate
                  Layout.fillWidth: true
                  height: 33
                  placeholderText: qsTr("Text Field")
                  text:Qt.formatDate(cal.selectedDate, "dd-MM-yyyy")
              }

              Button {
                  id: button
                  anchors.left: textDate.right
                  width: 25
                  height: 29
                  Image {
                      anchors.horizontalCenter: parent.horizontalCenter
                      id: img
                      width: 36
                          height: 44
                          source: "qrc:///qml/images/white/add.png"
                  }
                  onClicked:{
                      calendarPopup.open()
                  }
              }

          }

          RowLayout {
          Layout.fillWidth: true

              TextField {
                  id: textDate2
                  Layout.fillWidth: true
                  height: 33
                  placeholderText: qsTr("Text Field")
                  text:Qt.formatDate(cal2.selectedDate, "dd-MM-yyyy")
              }

              Button {
                  id: button2
                  anchors.left: textDate2.right
                  width: 25
                  height: 29
                  Image {
                      id: img2
                      anchors.horizontalCenter: parent.horizontalCenter
                      width: 36
                      height: 44
                      source: "qrc:///qml/images/white/add.png"
                  }
                  onClicked:{
                      calendarPopup2.open()
                  }
              }

          }

          RowLayout {

              Button {
                 id: proccessConsumptions
                 Layout.fillWidth: true
                 spacing: 1
                 text: "Consultar"
                 onClicked: Controller.getConsumptions(client.data.sensorID, textDate.text, textDate2.text)
              }

          }

          RowLayout {

              Pane {
                  id: hPane
                  Layout.fillWidth: true
                  Layout.fillHeight: true

                  ListModel {
                     id: listModel
                  }

                  TableView {
                      anchors.fill: parent
//                      backgroundVisible: false
                      Layout.fillHeight: true
                      Layout.fillWidth: true
                      alternatingRowColors: true

                      TableViewColumn {
                          role: "date"
                          title: "Data"
                          width: 150
                      }

                      TableViewColumn {
                          role: "measure"
                          title: "Medida"
                          width: 150
                      }

                      model: listModel
                  }


                  Connections {
                     target: Controller
                     onGetConsumption: {

                        var jsonObject = JSON.parse(consumptions);
                        listModel.clear()

                        Object.keys(jsonObject).forEach(function(k) {

                          jsonObject[k].forEach(function(c){

                           listModel.append({
                             date: new Date(c.date).toLocaleString(Qt.locale("pt_BR"), "dd/MM/yy HH:mm:ss"),
                             measure: c.measure
                           })

                          })

                        })

//                       console.log(Object.keys(jsonObject))
                     }
                  }
              }


          }

         } // col layout


    Popup {
        id: calendarPopup
        modal: true
        focus: true
        x: (window.width - width) / 2
        y: window.height / 6
        width: Math.min(window.width, window.height) / 3 * 2
        contentHeight: calendarColumn.height
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

        Column {
            id: calendarColumn
            spacing: 20

            Label {
                text: "Selecione uma data"
                font.bold: true
            }

            Calendar {
                 id: cal
                 Layout.fillWidth: true
                 Layout.fillHeight: true
                 minimumDate:  new Date(2018, 0, 1)
                 maximumDate: new Date()
                 selectedDate: new Date()
                 onClicked:  {
                      textDate.text = Qt.formatDate(cal.selectedDate, "dd-MM-yyyy");
                      calendarPopup.close()
                 }
            }
        }
    }

    Popup {
            id: calendarPopup2
            modal: true
            focus: true
            x: (window.width - width) / 2
            y: window.height / 6
            width: Math.min(window.width, window.height) / 3 * 2
            contentHeight: calendarColumn.height
            closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

            Column {
                id: calendarColumn2
                spacing: 20

                Label {
                    text: "Selecione uma data"
                    font.bold: true
                }

                Calendar {
                     id: cal2
                     Layout.fillWidth: true
                     Layout.fillHeight: true
                     selectedDate: new Date()
                     onClicked:  {
                          textDate2.text = Qt.formatDate(cal2.selectedDate, "dd-MM-yyyy");
                          calendarPopup2.close()
                     }
                }
            }
        }


  }

