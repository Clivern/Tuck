// Tuck docs: theme picker, copy buttons, playground, filtering, and keyboard nav.
(function () {
  var root = document.documentElement;
  var themes = ["blue", "slate", "emerald", "dark"];

  // Theme picker. The choice is stored under the same key Nut uses.
  var select = document.querySelector("[data-theme-select]");
  if (select) {
    var current = themes.find(function (t) {
      return root.classList.contains("theme-" + t);
    });
    select.value = current || "default";

    select.addEventListener("change", function () {
      themes.forEach(function (t) {
        root.classList.remove("theme-" + t);
      });

      if (select.value !== "default") {
        root.classList.add("theme-" + select.value);
      }

      try {
        localStorage.setItem("_ftheme", select.value);
      } catch (e) {}
    });
  }

  // Copy buttons.
  function copy(text, button) {
    var done = function () {
      button.classList.add("copied");
      setTimeout(function () {
        button.classList.remove("copied");
      }, 1200);
    };

    if (navigator.clipboard) {
      navigator.clipboard.writeText(text).then(done, function () {});
    }
  }

  document.addEventListener("click", function (e) {
    var button = e.target.closest("[data-copy], [data-copy-text]");
    if (!button) return;

    var text = button.dataset.copyText;
    if (text === undefined) {
      text = button.closest(".code").querySelector("pre").textContent;
    }

    copy(text, button);
  });

  // Open the program in the Go Playground. go.dev does not accept cross-origin share
  // requests, so copy the source to the clipboard and open an empty playground to
  // paste it into.
  var play = document.querySelector("[data-play]");
  if (play) {
    play.addEventListener("click", function () {
      var code = play.parentNode.querySelector("[data-copy-text]").dataset.copyText;
      var tab = window.open("https://go.dev/play/", "_blank");
      if (navigator.clipboard) {
        navigator.clipboard.writeText(code).then(function () {
          play.classList.add("copied");
          setTimeout(function () {
            play.classList.remove("copied");
          }, 1200);
        });
      }
      if (!tab) window.location.href = "https://go.dev/play/";
    });
  }

  // Filter examples in the sidebar or on the home page.
  var filter = document.querySelector("[data-filter]");
  if (filter) {
    filter.addEventListener("input", function () {
      var q = filter.value.trim().toLowerCase();
      var groups = document.querySelectorAll(".side-part, [data-part]");

      groups.forEach(function (group) {
        var shown = 0;
        group.querySelectorAll(".side-link, .lesson-card").forEach(function (link) {
          var hit = !q || link.textContent.toLowerCase().indexOf(q) !== -1;
          link.hidden = !hit;
          if (hit) shown++;
        });

        group.hidden = shown === 0;
      });
    });

    filter.addEventListener("keydown", function (e) {
      if (e.key === "Enter") {
        var first = document.querySelector(".side-link:not([hidden]), .lesson-card:not([hidden])");
        if (first) location.href = first.href;
      } else if (e.key === "Escape") {
        filter.value = "";
        filter.dispatchEvent(new Event("input"));
        filter.blur();
      }
    });
  }

  // Mobile sidebar.
  var menu = document.querySelector("[data-menu]");
  var sidebar = document.querySelector("[data-sidebar]");
  if (menu && sidebar) {
    menu.addEventListener("click", function () {
      sidebar.classList.toggle("open");
    });
  } else if (menu) {
    menu.hidden = true;
  }

  // Keep the active sidebar link in view.
  var active = document.querySelector(".side-link.active");
  if (active && sidebar) {
    sidebar.scrollTop = active.offsetTop - sidebar.clientHeight / 3;
  }

  // Highlight the section being read in "On this page".
  var tocLinks = document.querySelectorAll(".toc a");
  if (tocLinks.length && "IntersectionObserver" in window) {
    var byId = {};
    tocLinks.forEach(function (a) {
      byId[a.getAttribute("href").slice(1)] = a;
    });

    var observer = new IntersectionObserver(
      function (entries) {
        entries.forEach(function (entry) {
          if (!entry.isIntersecting) return;
          tocLinks.forEach(function (a) {
            a.classList.remove("active");
          });
          byId[entry.target.id].classList.add("active");
        });
      },
      { rootMargin: "-80px 0px -65% 0px" }
    );

    Object.keys(byId).forEach(function (id) {
      var el = document.getElementById(id);
      if (el) observer.observe(el);
    });
  }

  // Keyboard: "/" focuses the filter, left and right arrows move between examples.
  document.addEventListener("keydown", function (e) {
    var tag = (e.target.tagName || "").toLowerCase();
    if (tag === "input" || tag === "select" || tag === "textarea") return;
    if (e.metaKey || e.ctrlKey || e.altKey) return;

    if (e.key === "/" && filter) {
      e.preventDefault();
      filter.focus();
    } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
      var link = document.querySelector(e.key === "ArrowLeft" ? "[data-prev]" : "[data-next]");
      if (link) location.href = link.href;
    }
  });
})();
