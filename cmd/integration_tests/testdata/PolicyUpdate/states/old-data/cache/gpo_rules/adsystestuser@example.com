- id: '{5EC4DF8F-FF4E-41DE-846B-52AA6FFAF242}'
  name: GPO1 for current User
  rules:
      dconf:
        - key: org/gnome/shell/favorite-apps
          value: |
              'libreoffice-writer.desktop'
              'snap-store_ubuntu-software.desktop'
              'yelp.desktop
          disabled: false
          meta: as
        - key: org/gnome/desktop/media-handling/automount
          value: true
          disabled: false
          meta: b
- id: '{073AA7FC-5C1A-4A12-9AFC-42EC9C5CAF04}'
  name: GPO2 for current User
  rules:
      dconf:
        - key: org/gnome/desktop/media-handling/automount
          value: ""
          disabled: true
          meta: ""
- id: '{75545F76-DEC2-4ADA-B7B8-D5209FD48727}'
  name: GPO for Integration Test User
  rules:
      dconf:
        - key: org/gnome/shell/old/old-data
          value: something
          disabled: false
          meta: s
        - key: org/gnome/desktop/background/picture-options
          value: none
          disabled: false
          meta: s
        - key: org/gnome/desktop/background/picture-uri
          value: file:///usr/share/backgrounds/ubuntu.png
          disabled: false
          meta: s
        - key: org/gnome/shell/favorite-apps
          value: |4
               'firefox.desktop'
              'thunderbird.desktop'
              'org.gnome.Nautilus.desktop'
          disabled: false
          meta: as
- id: '{31B2F340-016D-11D2-945F-00C04FB984F9}'
  name: Default Domain Policy
  rules: {}
