import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 1.2
import QtQuick.Controls 2.0
import QtQuick.Controls.Material 2.0
import QtQml 2.0


  Pane {
      id: root
      anchors.fill: parent

      ColumnLayout {
          anchors.right: parent.right
          anchors.left: parent.left

          RowLayout {

              TextField {
                  id: textDate
                  width: 300
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

              TextField {
                  id: textDate2
                  width: 300
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
                  width: window.width
                  Layout.fillHeight: true

                  ListModel {
                     id: listModel
                  }

                  TableView {
                      backgroundVisible: false
                      Layout.fillHeight: true
                      alternatingRowColors: false
                      width: 350

                      TableViewColumn {
                          role: "date"
                          title: "Data"
                          width: 200
                      }

                      TableViewColumn {
                          role: "measure"
                          title: "Medida"
                          width: 150
                      }

                      model: listModel
                  }

//                      ListView {
//                          id: listView
//                          anchors.fill: parent
//                          model: listModel
//                          width: parent.width
//
//                          delegate: Item {
//                              id: delegateItem
//                              width: parent.width; height: 100
//
//
//                              Text {
//                                  id: dateItemTxt
//                                  height: parent.height; width: 150
//                                  anchors.left: parent.left
//                                  // delegate can directly use ListElement role name
//                                  text: new Date(model.date).toLocaleString(Qt.locale("pt_BR"), "dd/MM/yy HH:mm:ss")
//                              }
//
//
//                              Text {
//                                  id: itexItem
//                                  anchors.left: dateItemTxt.right
//                                  anchors.leftMargin: 20
////                                  anchors.verticalCenter: parent.verticalCenter
////                                  font.pixelSize: 40
//                                  // delegate can directly use ListElement role name
//                                  text: model.measure
//                              }
//                          }
//
//
//                      }

                  Connections {
                     target: Controller
                     onGetConsumption: {

                       var jsonObject = JSON.parse(consumptions);

                       Object.keys(jsonObject).forEach(function(k) {

                         jsonObject[k].forEach(function(c){
                           listModel.append({
                             date: new Date(c.date).toLocaleString(Qt.locale("pt_BR"), "dd/MM/yy HH:mm:ss"),
                             measure: c.measure
                           })

                           console.log(listModel.count)
                         })

                       })

                       console.log(Object.keys(jsonObject))
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

