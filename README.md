# gowinx

Simple library to handle ux components on windows with go.  

The framework allows to click on an icon from the notification area. The icon should be identify by its name (text). After this,
when the menu holding the functionality of the icon is displayed the framework allows to click on any specific button of the menu. Menu options shoud be identify by its name.  

The framework, calculate the coordenates of the elemens to click on them through the use of the mouse, this is intetional to mimic as much as possible the user experience (instead of relying on windows messaging mechanism)*

*This option can be used but visualizations may differ from what expected.

# Windows/Controls

Window element can be considered controls, they offer handlers to interact with them. In case of notification area
there is a parent window class Shell_TrayWnd which contains several window / control objects, some of them:

* Toolbarwindows32, icons on notification area are splitted across two toolbars: 1 holding visible notification area icons and other with hidden icons.
* TrayButton which offers access to Action Center
* TrayClockWClass which offers access to time/date functionality

## Controls relationships

![win32ux](docs/diagrams/win32ux.jpg?raw=true)

# Bibliography

* [win32 ux guide](https://docs.microsoft.com/en-us/windows/win32/uxguide/guidelines)   
* [win32 api](https://docs.microsoft.com/en-us/windows/win32/api/_base/) 
* [notification area](https://docs.microsoft.com/en-us/windows/win32/shell/notification-area) 
