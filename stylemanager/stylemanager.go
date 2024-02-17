package stylemanager

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// Style represents CSS properties for an element.
type Style map[string]string

// Keyframes represents CSS keyframes for an animation.
type Keyframes map[string]Style

// CompositeStyle represents a collection of styles.
type CompositeStyle struct {
	Default       Style
	PseudoClasses map[string]Style
}

// StyleSheet represents a collection of styles mapped to class names.
type StyleSheet map[string]Style

// StyleManager manages styles and generates CSS classes.
type StyleManager struct {
	styles         StyleSheet
	compositeCache map[string]CompositeStyle
	animationCache map[string]Keyframes
}

// NewStyleManager creates a new instance of StyleManager.
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:         make(StyleSheet),
		animationCache: make(map[string]Keyframes),
		compositeCache: make(map[string]CompositeStyle),
	}
}

// AddStyle adds a new style to the manager and returns a class name.
func (sm *StyleManager) AddStyle(style Style) string {
	// Convert style to a string to generate hash.
	styleStr := fmt.Sprintf("%v", style)
	hash := sha1.Sum([]byte(styleStr))
	className := fmt.Sprintf("cls_%x", hash[:5]) // Use first 5 bytes of hash for class name.

	if _, exists := sm.styles[className]; !exists {
		sm.styles[className] = style
	}

	return className
}

// AddAnimation adds a new keyframes animation to the manager and returns an animation name.
func (sm *StyleManager) AddAnimation(keyframes Keyframes) string {
	// Convert keyframes to a string to generate hash.
	keyframesStr := fmt.Sprintf("%v", keyframes)
	hash := sha1.Sum([]byte(keyframesStr))
	animationName := fmt.Sprintf("anim_%x", hash[:5]) // Use first 5 bytes of hash for animation name.

	if _, exists := sm.animationCache[animationName]; !exists {
		sm.animationCache[animationName] = keyframes
	}

	return animationName
}

func (sm *StyleManager) AddCompositeStyle(composite CompositeStyle) string {
	// Convert composite to a string to generate hash.
	compositeStr := fmt.Sprintf("%v", composite)
	hash := sha1.Sum([]byte(compositeStr))
	className := fmt.Sprintf("cls_%x", hash[:5]) // Use first 5 bytes of hash for class name.

	if _, exists := sm.compositeCache[className]; !exists {
		sm.compositeCache[className] = composite
	}

	return className
}

// GenerateCSS generates the CSS string for all styles managed by StyleManager.
func (sm *StyleManager) GenerateCSS() string {
	var builder strings.Builder

	for className, style := range sm.styles {
		builder.WriteString(fmt.Sprintf(".%s { ", className))
		for prop, value := range style {
			builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
		}
		builder.WriteString("} ")
	}

	for animationName, keyframes := range sm.animationCache {
		builder.WriteString(fmt.Sprintf("@keyframes %s { ", animationName))
		for key, style := range keyframes {
			builder.WriteString(fmt.Sprintf("%s { ", key))
			for prop, value := range style {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
			}
			builder.WriteString("} ")
		}
		builder.WriteString("} ")
	}

	for className, composite := range sm.compositeCache {
		builder.WriteString(fmt.Sprintf(".%s { ", className))
		for prop, value := range composite.Default {
			builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
		}
		builder.WriteString("} ")

		for pseudoClass, style := range composite.PseudoClasses {
			// Check if pseudoClass already starts with a colon
			if strings.HasPrefix(pseudoClass, ":") {
				builder.WriteString(fmt.Sprintf(".%s%s { ", className, pseudoClass))
			} else {
				builder.WriteString(fmt.Sprintf(".%s:%s { ", className, pseudoClass))
			}
			for prop, value := range style {
				builder.WriteString(fmt.Sprintf("%s: %s; ", prop, value))
			}
			builder.WriteString("} ")
		}
	}

	return builder.String()
}
