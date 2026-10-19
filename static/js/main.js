(function () {
  "use strict";

  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  var finePointer = window.matchMedia("(pointer: fine)").matches;

  /* Mobile nav */
  var toggle = document.getElementById("nav-toggle");
  var nav = document.getElementById("nav");
  if (toggle && nav) {
    toggle.addEventListener("click", function () {
      nav.classList.toggle("open");
    });
    nav.querySelectorAll("a").forEach(function (link) {
      link.addEventListener("click", function () {
        nav.classList.remove("open");
      });
    });
  }

  /* Cursor spotlight */
  var spotlight = document.getElementById("spotlight");
  if (spotlight && finePointer && !reduceMotion) {
    document.body.classList.add("has-pointer");
    var sx = 0, sy = 0, tx = 0, ty = 0;
    window.addEventListener("pointermove", function (e) {
      tx = e.clientX;
      ty = e.clientY;
    }, { passive: true });
    (function loopSpot() {
      sx += (tx - sx) * 0.12;
      sy += (ty - sy) * 0.12;
      spotlight.style.transform = "translate(" + (sx - 300) + "px," + (sy - 300) + "px)";
      requestAnimationFrame(loopSpot);
    })();
  }

  /* Parallax layers */
  var parallaxEls = document.querySelectorAll("[data-parallax]");
  if (parallaxEls.length && !reduceMotion) {
    var pItems = [];
    parallaxEls.forEach(function (el) {
      pItems.push({ el: el, factor: parseFloat(el.getAttribute("data-parallax")) || 0.05 });
    });
    var py = 0, pty = 0;
    window.addEventListener("scroll", function () {
      pty = window.scrollY || 0;
    }, { passive: true });
    (function loopParallax() {
      py += (pty - py) * 0.08;
      pItems.forEach(function (item) {
        item.el.style.transform = "translate3d(0," + (py * item.factor * -1).toFixed(2) + "px,0)";
      });
      requestAnimationFrame(loopParallax);
    })();
  }

  /* 3D tilt + glare */
  if (finePointer && !reduceMotion) {
    document.querySelectorAll("[data-tilt]").forEach(function (el) {
      var max = parseFloat(el.getAttribute("data-tilt-max")) || 8;
      var glare = el.hasAttribute("data-tilt-glare");
      var rect = null;
      var currentX = 0, currentY = 0;
      var targetX = 0, targetY = 0;
      var glareX = 50, glareY = 50;
      var tGlareX = 50, tGlareY = 50;
      var active = false;
      var raf = null;

      function tick() {
        currentX += (targetX - currentX) * 0.12;
        currentY += (targetY - currentY) * 0.12;
        glareX += (tGlareX - glareX) * 0.15;
        glareY += (tGlareY - glareY) * 0.15;
        el.style.transform =
          "perspective(1000px) rotateX(" + currentY.toFixed(3) + "deg) rotateY(" + currentX.toFixed(3) + "deg) scale3d(1.015,1.015,1.015)";
        if (glare) {
          el.style.setProperty("--glare-x", glareX.toFixed(2) + "%");
          el.style.setProperty("--glare-y", glareY.toFixed(2) + "%");
        }
        var dx = Math.abs(targetX - currentX) + Math.abs(targetY - currentY) +
                 Math.abs(tGlareX - glareX) + Math.abs(tGlareY - glareY);
        if (dx > 0.02) {
          raf = requestAnimationFrame(tick);
        } else {
          raf = null;
        }
      }

      function ensureTick() {
        if (!raf) raf = requestAnimationFrame(tick);
      }

      el.addEventListener("pointerenter", function () {
        rect = el.getBoundingClientRect();
        active = true;
      });

      el.addEventListener("pointermove", function (e) {
        if (!rect) rect = el.getBoundingClientRect();
        var px = (e.clientX - rect.left) / rect.width;
        var py = (e.clientY - rect.top) / rect.height;
        targetX = (px - 0.5) * max * 2;
        targetY = (0.5 - py) * max * 2;
        tGlareX = px * 100;
        tGlareY = py * 100;
        ensureTick();
      });

      el.addEventListener("pointerleave", function () {
        active = false;
        rect = null;
        targetX = 0;
        targetY = 0;
        tGlareX = 50;
        tGlareY = 50;
        ensureTick();
      });
    });
  }

  /* Magnetic buttons */
  if (finePointer && !reduceMotion) {
    document.querySelectorAll("[data-magnetic]").forEach(function (el) {
      var strength = 18;
      el.addEventListener("pointermove", function (e) {
        var r = el.getBoundingClientRect();
        var x = e.clientX - (r.left + r.width / 2);
        var y = e.clientY - (r.top + r.height / 2);
        el.style.transform = "translate(" + (x / r.width) * strength + "px," + (y / r.height) * strength + "px)";
      });
      el.addEventListener("pointerleave", function () {
        el.style.transition = "transform 0.35s cubic-bezier(0.22,1,0.36,1)";
        el.style.transform = "translate(0,0)";
        setTimeout(function () { el.style.transition = ""; }, 360);
      });
    });
  }

  /* Scroll reveal */
  var revealEls = document.querySelectorAll("[data-reveal]");
  if (revealEls.length) {
    if (reduceMotion || !("IntersectionObserver" in window)) {
      revealEls.forEach(function (el) { el.classList.add("is-visible"); });
    } else {
      var io = new IntersectionObserver(function (entries) {
        entries.forEach(function (entry) {
          if (entry.isIntersecting) {
            entry.target.classList.add("is-visible");
            io.unobserve(entry.target);
          }
        });
      }, { threshold: 0.12, rootMargin: "0px 0px -40px 0px" });
      revealEls.forEach(function (el) { io.observe(el); });
    }
  }

  /* Contact form */
  var form = document.getElementById("contact-form");
  var status = document.getElementById("form-status");
  if (form) {
    form.addEventListener("submit", async function (e) {
      e.preventDefault();
      status.textContent = "";
      status.className = "form-status";

      var data = new FormData(form);
      try {
        var res = await fetch("/contact", {
          method: "POST",
          body: data,
          headers: { "X-Requested-With": "XMLHttpRequest", Accept: "application/json" },
        });
        if (res.ok) {
          status.textContent = "Thanks — your message has been sent.";
          status.classList.add("ok");
          form.reset();
        } else {
          var text = await res.text();
          status.textContent = text || "Something went wrong. Please try again.";
          status.classList.add("err");
        }
      } catch (err) {
        status.textContent = "Network error. Please try again.";
        status.classList.add("err");
      }
    });
  }
})();
