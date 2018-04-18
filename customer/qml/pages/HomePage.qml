import QtQuick 2.6
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0

Pane {
    id: homePane
    anchors.fill: parent


    property string counterCache: "0"

    ColumnLayout {
        anchors.rightMargin: 0
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 29
        anchors.fill: parent

        Label {
            text: qsTr("Seja Bem Vindo %1").arg(client.data.name)
            Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
            horizontalAlignment: Qt.AlignHCenter
            font.pixelSize: 15
//            wrapMode: Label.wrapMode
        }

        Column {
            width: homePane.availableWidth / 2
            height: homePane.availableHeight / 2
            anchors.centerIn: parent
            anchors.verticalCenterOffset: -50
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter

            Rectangle {
                id: rectCount
                width: 200
                height: 200
                color: "#00000000"
                border.color: "#0591c8"
                radius: width * 0.5


                Label {
                    id: countLabel
                    text: "0"
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.horizontalCenter: parent.horizontalCenter
                    font.pixelSize: 35
                    width: parent.width
                    horizontalAlignment: Qt.AlignHCenter
                    anchors.margins: 10

                    Component.onCompleted: {
                      if (counterCache == "0") {
                        counterCache = Controller.getConsumptionValue(client.data.sensorID)
                        countLabel.text = counterCache
                      }

                    }

                    Text {
                        anchors.top: parent.bottom
                        anchors.horizontalCenter: parent.horizontalCenter
                        font.pixelSize: 10
                        text: qsTr("Seu consumo hoje")
                        color: Material.color(Material.Blue)
                    }
                }


            }


        }


    }

}

