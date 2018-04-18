import QtQuick 2.7
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

ApplicationWindow {
  id: window
  width: 360
  height: 400
  visible: true
  title: "SMCRA - Sensor"

  readonly property int itemWidth: Math.max(box.implicitWidth, Math.min(box.implicitWidth * 2, pane.availableWidth / 3))

  StackView {
    id: stackView
    anchors.fill: parent

    initialItem: Pane {
      id: pane
      x: 0
      y: 0

      Label {
        text: qsTr("")
        anchors.horizontalCenter: parent.horizontalCenter
        Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
        horizontalAlignment: Qt.AlignHCenter
        font.pixelSize: 15
        wrapMode: Label.wrapMode

      }

      ColumnLayout {
        id: columnLayout1
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 29
        anchors.fill: parent


        Column {
          id: column1
          width: pane.availableWidth
          anchors.fill: parent
          Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

          Rectangle {
              id: rectangle1
              width: pane.availableWidth
              height: 100

              Text {
                id: title
                color: "black"
                font.pointSize: 16
                horizontalAlignment: Qt.AlignHCenter
                text: "0"
                style: Text.Normal
                font.bold: false
                anchors.horizontalCenter: parent.horizontalCenter
                anchors.verticalCenter: parent.verticalCenter

                Connections {
                  target: Controller
                  onUpdateCounter: title.text = data + "ml/s"
                }

              }
          }


          SpinBox {
              id: box
              from: 0
              value: 0
              to: 20
              width: itemWidth
              anchors.top: rectangle1.bottom
              anchors.topMargin: 20
              anchors.horizontalCenter: parent.horizontalCenter
              onValueChanged: Controller.updateFlowRate(box.value)
          }

//          Button {
//            text: "Button Text"
//            spacing: 3
//            checkable: true
//            Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
//            anchors.verticalCenter: parent.verticalCenter
//            anchors.horizontalCenter: parent.horizontalCenter
//
//            onClicked: console.log("hello from qml")
//          }

        }

      }


    }
  }
}

