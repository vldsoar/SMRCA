import QtQuick 2.6
import QtQuick.Controls 2.0

Page {
    id: page
    anchors.fill: parent
    property var pages: ["qrc:///qml/pages/HomePage.qml", "qrc:///qml/pages/HistoryPage.qml"]

    SwipeView {
        id: swipeView
        anchors.fill: parent
        currentIndex: tabBar.currentIndex

        Repeater {
            model: 2

            Loader {
                id: loader
                active: true
                source: pages[swipeView.currentIndex]
            }
        }

    }

    footer: TabBar {
        id: tabBar
        currentIndex: swipeView.currentIndex

        TabButton {
            text: "Home"
        }
        TabButton {
            text: "Histórico"
        }
    }
}