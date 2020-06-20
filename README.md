[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: `framekit` desktop, mobile, and display GUI Framework
**URL** [multiverse-os.org](https://multiverse-os.org)

Multiverse OS development requires maintaing a variety of basic application
frameworks that are designed for ease-of-use, simplicity, similarity-- by
implementing what works in existing popular frameworks; but stritly following
the Multiverse OS design guidelines, and maintaing a strict adherence to a
security first approach. 

Currently in active flux, there are three implementation being experimented
with actively by different developers. 

    1) A webkit implementation with Javascript completely disabled; currently
    working out how to implement it using the incredibly reduced `WPE webkit`
    for embedded low-power consumption computer devices, and designed from the
    ground up for increased security. 

    2) One utilizes ChromeDP, embedded chromium (soon a lightweight chromium), and
    associated tools to provide a jailed, UI without network. Possible support
    for very limited, signed javascript. 

    3) QT WebEngine based solution, however with QT, we can just use QML, and
    avoid the requirement of web-engines, and their security issues; but we
    can't use our web-framework repeatedly. 
      * Similary, GTK allows for a CSS styling that can provide wonderful
        results without requiring full browser engine; with the primary problem
        is that we want to migrate away from Gnome specifically because it uses
        Javascript for its UI, and it uses a relatively not-battle tested
        javascript implementation. 

    4) Warp Engine, a simple web engine that has already begun development and
    shows promise for a variety of reasons, and limitations associated with
    using an off the shelf web-engine. 

    5) Servo Engine, experiments using this have been quite nice, and this is
    definitely a possible solution. At the very least, we may provide a generic
    system for working with these web engines, even if we do not actively use
    one for your default GUI framework; because of how common they have become,
    so that we can provide a minimum security standard; that all other
    applicaitons using this style UI in Multiverse OS mus comply with. 

#### Contribution
The best way to contribute to framekit currently is to either implement desired
features, request features you would use, like our servo experiment; and begin
using it, and improve it, and notify us of what changes need to be made. 
